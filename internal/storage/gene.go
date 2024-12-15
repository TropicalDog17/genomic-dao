package storage

import (
	"encoding/hex"

	"gorm.io/gorm"
)

// GeneData represents the structure to hold gene data and associated information.
type GeneData struct {
	gorm.Model
	FileID        string
	UserID        uint32
	EncryptedData []byte
	DataHash      []byte
	Signature     []byte
}

type GenDataRepository interface {
	StoreGeneData(userID uint32, encryptedData []byte, signatureBytes []byte, hashBytes []byte) (string, error)
	RetrieveGeneData(fileID string) ([]byte, error)
}

type genDataRepository struct {
	db *gorm.DB
}

func NewGenDataRepository(db *gorm.DB) GenDataRepository {
	return &genDataRepository{db}
}

func (r *genDataRepository) StoreGeneData(userID uint32, encryptedData []byte, signatureBytes []byte, hashBytes []byte) (string, error) {
	fileID := hex.EncodeToString(hashBytes[:8])
	geneData := GeneData{
		FileID:        fileID,
		UserID:        userID,
		EncryptedData: encryptedData,
		DataHash:      hashBytes,
		Signature:     signatureBytes,
	}

	if err := r.db.Create(&geneData).Error; err != nil {
		return "", err
	}

	return geneData.FileID, nil
}

func (r *genDataRepository) RetrieveGeneData(fileID string) ([]byte, error) {
	var geneData GeneData
	if err := r.db.Where("file_id = ?", fileID).First(&geneData).Error; err != nil {
		return nil, err
	}

	return geneData.EncryptedData, nil
}
