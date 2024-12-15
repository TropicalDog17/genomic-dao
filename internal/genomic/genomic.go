package genomic

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/TropicalDog17/genomic-dao-service/internal/auth"
	"github.com/TropicalDog17/genomic-dao-service/internal/onchain"
	"github.com/TropicalDog17/genomic-dao-service/internal/storage"
	"github.com/TropicalDog17/genomic-dao-service/internal/tee"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

type UploadResult struct {
	SessionID string
	FileID    string
	Message   string
}

type genomicService struct {
	teeService             tee.TeeService
	geneDataStorageService storage.GeneDataStorageService
	authService            auth.AuthService
	onchainService         onchain.OnchainService
}

type GenomicService interface {
	ProcessAndUploadGenomicData(genomicData []byte, pubkey string, privateKey *ecdsa.PrivateKey) (*UploadResult, error)
	RetrieveGenomicData(fileID string, privKey *ecdsa.PrivateKey) ([]byte, error)
}

func NewGenomicService(
	teeService tee.TeeService,
	geneDataStorageService storage.GeneDataStorageService,
	authService auth.AuthService,
	onchainService onchain.OnchainService,
) GenomicService {
	return &genomicService{
		teeService:             teeService,
		geneDataStorageService: geneDataStorageService,
		authService:            authService,
		onchainService:         onchainService,
	}
}

func (s *genomicService) RetrieveGenomicData(fileID string, privKey *ecdsa.PrivateKey) ([]byte, error) {

	// Retrieve genomic data from storage
	genomicData, err := s.geneDataStorageService.RetrieveGeneData(fileID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve genomic data: %w", err)
	}

	decryptedGenomicData, err := s.teeService.DecryptData(genomicData, privKey)

	if err != nil {
		return nil, fmt.Errorf("failed to decrypt genomic data: %w", err)
	}

	return decryptedGenomicData, nil
}

func (s *genomicService) ProcessAndUploadGenomicData(genomicData []byte, pubkey string, privateKey *ecdsa.PrivateKey) (*UploadResult, error) {
	// Authenticate user
	user, err := s.authService.Authenticate(pubkey)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate user: %w", err)
	}

	// Process genomic data in TEE and get risk score at the same time
	processedData, riskScore, err := s.teeService.ProcessAndEncrypt(genomicData, &privateKey.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to process gene data: %w", err)
	}

	// Sign the processed data
	hash, signature, err := s.signEncryptedGeneData(privateKey, processedData)
	if err != nil {
		return nil, fmt.Errorf("failed to sign gene data: %w", err)
	}

	// Store the data
	fileID, err := s.geneDataStorageService.StoreGeneData(user.ID, processedData, hash, signature)
	if err != nil {
		return nil, fmt.Errorf("failed to store gene data: %w", err)
	}

	// Upload to blockchain
	sessionID, err := s.onchainService.UploadData(fileID)
	if err != nil {
		return nil, fmt.Errorf("failed to upload data to blockchain: %w", err)
	}

	// Confirm upload
	docID := uuid.New().String()
	err = s.onchainService.ConfirmUpload(docID, hex.EncodeToString(hash), "0x1234", sessionID, riskScore)
	if err != nil {
		return nil, fmt.Errorf("failed to confirm upload: %w", err)
	}

	return &UploadResult{
		SessionID: sessionID,
		Message:   "Genomic data uploaded successfully",
		FileID:    fileID,
	}, nil
}

func (s *genomicService) signEncryptedGeneData(privateKey *ecdsa.PrivateKey, encryptedData []byte) ([]byte, []byte, error) {
	hashData := crypto.Keccak256Hash(encryptedData).Bytes()
	signature, err := crypto.Sign(hashData, privateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to sign data: %w", err)
	}
	return hashData, signature, nil
}
