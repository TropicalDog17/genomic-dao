package handler

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"

	"github.com/TropicalDog17/genomic-dao-service/internal/auth"
	"github.com/TropicalDog17/genomic-dao-service/internal/onchain"
	"github.com/TropicalDog17/genomic-dao-service/internal/storage"
	"github.com/TropicalDog17/genomic-dao-service/internal/tee"
	genomicCrypto "github.com/TropicalDog17/genomic-dao-service/pkg/crypto"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type genomicHandler struct {
	teeService             *tee.TeeService
	geneDataStorageService storage.GeneDataStorageService
	authService            *auth.AuthService
	onchainService         *onchain.OnchainService
}

type GenomicHandler interface {
	UploadGenomicData(c *gin.Context)
}

func NewGenomicHandler(teeService *tee.TeeService, geneDataStorageService storage.GeneDataStorageService, authService *auth.AuthService, onchainService *onchain.OnchainService) GenomicHandler {
	return &genomicHandler{
		teeService:             teeService,
		geneDataStorageService: geneDataStorageService,
		authService:            authService,
		onchainService:         onchainService,
	}
}

// @Summary Upload genomic data for processing
// @Description Processes genomic data in TEE, encrypts it, calculates risk score, and stores on blockchain
// @Tags genomic
// @Accept multipart/form-data
// @Produce json
// @Param genomicData formData string true "Raw genomic data to be processed"
// @Param pubkey formData string true "User's public key for authentication"
// @Success 200 {object} map[string]string "Returns transaction hash"
// @Failure 500 {object} map[string]string "Internal server error with error message"
// @Router /upload [post]
func (h *genomicHandler) UploadGenomicData(c *gin.Context) {
	genomicData := c.PostForm("genomicData")
	pubkey := c.PostForm("pubkey")

	// auth the user use the public key for demo simplicity
	user, err := h.authService.UserRepository.FindByPubkey(pubkey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Upload the genomic data
	data, err := h.teeService.ProcessGeneData([]byte(genomicData), []byte(pubkey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// get user private key from env for simplicity
	privateKey, _, _, err := genomicCrypto.DeriveEcdsaKeyPairAndEthAddress(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Sign the genomic data
	hash, signature, err := signEncryptedGeneData(privateKey, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Genomic data signed successfully")

	// Store the encrypted genomic data and other details in the storage
	fileID, err := h.geneDataStorageService.StoreGeneData(user.ID, data, hash, signature)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	riskScore, err := h.teeService.CalculateRiskScore(data, []byte(pubkey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Risk score calculated successfully", riskScore)

	sessionId, err := h.onchainService.UploadData(fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Confirm the genomic data upload
	docId := uuid.New().String()
	err = h.onchainService.ConfirmUpload(docId, hex.EncodeToString(hash), "0x1234", sessionId, riskScore)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"sessionId": sessionId, "message": "Genomic data uploaded successfully"})

}

// signEncryptedGeneData generates a hash of the encrypted data and then signs it using the provided private key.
// It returns the hash and the corresponding digital signature.
func signEncryptedGeneData(privateKey *ecdsa.PrivateKey, encryptedData []byte) ([]byte, []byte, error) {
	// Generate the hash of the encrypted data
	hashData := crypto.Keccak256Hash(encryptedData).Bytes()

	// Sign the hash using the provided private key
	signature, err := crypto.Sign(hashData, privateKey)
	if err != nil {
		return nil, nil, err
	}

	return hashData, signature, nil
}
