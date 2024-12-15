package server

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/TropicalDog17/genomic-dao-service/internal/auth"
	"github.com/TropicalDog17/genomic-dao-service/internal/handler"
	"github.com/TropicalDog17/genomic-dao-service/internal/onchain"
	"github.com/TropicalDog17/genomic-dao-service/internal/storage"
	"github.com/TropicalDog17/genomic-dao-service/internal/tee"
	genomicCrypto "github.com/TropicalDog17/genomic-dao-service/pkg/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var startOnce sync.Once

func StartService(db *gorm.DB) {
	startOnce.Do(func() {
		r := route(db)
		err := r.Run(":8080")
		if err != nil {
			panic(err)
		}
		fmt.Println("Server started at :8080")
	})
}

// Routing by gin
func route(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		panic(err)
	}

	// Get the chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		panic(fmt.Errorf("failed to get chain ID: %v", err))

	}

	// get owner private key(BIP44) from env for simplicity
	privKey, _, _, err := genomicCrypto.DeriveEcdsaKeyPairAndEthAddress(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		panic(fmt.Errorf("failed to derive ECDSA key pair and Ethereum address: %v", err))
	}
	// Create a new transactor with chain ID
	opts, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		panic(fmt.Errorf("failed to create transactor: %v", err))
	}
	authService := auth.NewAuthService(auth.NewUserRepository(db))
	// storageService := storage.NewGenDataRepository(db)

	authHandler := handler.NewAuthHandler(authService)
	teeService := tee.NewTeeService()
	geneDataStorageService := storage.NewGeneDataStorageService(db)

	controllerContractAddress := common.HexToAddress(os.Getenv("CONTROLLER_ADDRESS"))
	_ = common.HexToAddress(os.Getenv("PCSP_ADDRESS"))

	onchainService := onchain.NewOnchainService(client, opts, controllerContractAddress)
	genomicHandler := handler.NewGenomicHandler(teeService, geneDataStorageService, authService, onchainService)
	pcspHandler := handler.NewPCSPHandler(onchain.NewPcspService(client, opts, common.HexToAddress(os.Getenv("PCSP_ADDRESS"))))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.GET("/pcsp/balance", pcspHandler.GetPCSPBalance)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Create a new user
	r.POST("/auth/register", authHandler.Register)

	// Upload genomic data
	r.POST("/upload", genomicHandler.UploadGenomicData)

	return r
}
