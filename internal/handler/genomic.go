package handler

import (
	"crypto/ecdsa"
	"net/http"
	"os"

	"github.com/TropicalDog17/genomic-dao-service/internal/auth"
	"github.com/TropicalDog17/genomic-dao-service/internal/genomic"
	"github.com/TropicalDog17/genomic-dao-service/internal/onchain"
	"github.com/TropicalDog17/genomic-dao-service/internal/storage"
	"github.com/TropicalDog17/genomic-dao-service/internal/tee"
	genomicCrypto "github.com/TropicalDog17/genomic-dao-service/pkg/crypto"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

// Response structures for Swagger documentation
type GenomicDataResponse struct {
	GenomicData string `json:"genomicData" example:"ATCG..."`
}

type UploadResponse struct {
	SessionID string `json:"sessionId" example:"sess_123"`
	Message   string `json:"message" example:"Upload successful"`
	FileID    string `json:"fileId" example:"file_123"`
}

type genomicHandler struct {
	teeService             tee.TeeService
	geneDataStorageService storage.GeneDataStorageService
	authService            auth.AuthService
	onchainService         onchain.OnchainService
	genomicService         genomic.GenomicService
}

type GenomicHandler interface {
	UploadGenomicData(c *gin.Context)
	RetrieveGenomicData(c *gin.Context)
}

func NewGenomicHandler(teeService tee.TeeService, geneDataStorageService storage.GeneDataStorageService, authService auth.AuthService, onchainService onchain.OnchainService) GenomicHandler {
	return &genomicHandler{
		teeService:             teeService,
		geneDataStorageService: geneDataStorageService,
		authService:            authService,
		onchainService:         onchainService,
		genomicService:         genomic.NewGenomicService(teeService, geneDataStorageService, authService, onchainService),
	}
}

// @Summary Retrieve genomic data
// @Description Retrieves genomic data from the blockchain
// @Tags genomic
// @Accept json
// @Produce json
// @Param fileID query string true "File ID of the genomic data"
// @Success 200 {object} GenomicDataResponse
// @Failure 500 {object} ErrorResponse
// @Router /retrieve [get]
func (h *genomicHandler) RetrieveGenomicData(c *gin.Context) {
	fileID := c.Query("fileID")
	privKey, _, _, err := genomicCrypto.DeriveEcdsaKeyPairAndEthAddress(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	genomicData, err := h.genomicService.RetrieveGenomicData(fileID, privKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GenomicDataResponse{
		GenomicData: string(genomicData),
	})
}

// @Summary Upload genomic data for processing
// @Description Processes genomic data in TEE, encrypts it, calculates risk score, and stores on blockchain
// @Tags genomic
// @Accept multipart/form-data
// @Produce json
// @Param genomicData formData string true "Raw genomic data to be processed"
// @Param pubkey formData string true "User's public key for authentication"
// @Success 200 {object} UploadResponse
// @Failure 500 {object} ErrorResponse
// @Router /upload [post]
func (h *genomicHandler) UploadGenomicData(c *gin.Context) {
	genomicData := c.PostForm("genomicData")
	pubkey := c.PostForm("pubkey")

	privateKey, _, _, err := genomicCrypto.DeriveEcdsaKeyPairAndEthAddress(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	result, err := h.genomicService.ProcessAndUploadGenomicData([]byte(genomicData), pubkey, privateKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, UploadResponse{
		SessionID: result.SessionID,
		Message:   result.Message,
		FileID:    result.FileID,
	})
}

func signEncryptedGeneData(privateKey *ecdsa.PrivateKey, encryptedData []byte) ([]byte, []byte, error) {
	hashData := crypto.Keccak256Hash(encryptedData).Bytes()
	signature, err := crypto.Sign(hashData, privateKey)
	if err != nil {
		return nil, nil, err
	}
	return hashData, signature, nil
}
