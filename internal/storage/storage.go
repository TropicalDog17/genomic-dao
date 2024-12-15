package storage

import "gorm.io/gorm"

type geneDataStorageService struct {
	geneDataRepository GenDataRepository
}

type GeneDataStorageService interface {
	StoreGeneData(userID uint32, encryptedData []byte, signatureBytes []byte, hashBytes []byte) (string, error)
	RetrieveGeneData(fileID string) ([]byte, error)
}

// NewGeneDataStorageService creates a new gene data storage service.
func NewGeneDataStorageService(db *gorm.DB) GeneDataStorageService {
	repo := NewGenDataRepository(db)
	return &geneDataStorageService{
		geneDataRepository: repo,
	}
}

// StoreGeneData stores gene data in the database.
func (s *geneDataStorageService) StoreGeneData(userID uint32, encryptedData []byte, signatureBytes []byte, hashBytes []byte) (string, error) {
	return s.geneDataRepository.StoreGeneData(userID, encryptedData, signatureBytes, hashBytes)
}

// RetrieveGeneData retrieves gene data from the database.
func (s *geneDataStorageService) RetrieveGeneData(fileID string) ([]byte, error) {
	return s.geneDataRepository.RetrieveGeneData(fileID)
}
