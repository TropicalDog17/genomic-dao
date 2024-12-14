package storage

import "gorm.io/gorm"

// GeneData represents the structure to hold gene data and associated information.
type GeneData struct {
	gorm.Model
	FileID        string
	UserID        uint64
	DataHash      []byte // Hash of the encrypted gene data.
	Signature     []byte // Signature of the hash.
	EncryptedData []byte // The encrypted gene data.
}

type GenDataRepository interface {
	StoreGeneData(userID uint64, encryptedData []byte, signatureBytes []byte, hashBytes []byte) (string, error)
	RetrieveGeneData(fileID string) ([]byte, error)
}

type genDataRepository struct {
	db *gorm.DB
}

func NewGenDataRepository(db *gorm.DB) GenDataRepository {
	return &genDataRepository{db}
}

func (r *genDataRepository) StoreGeneData(userID uint64, encryptedData []byte, signatureBytes []byte, hashBytes []byte) (string, error) {
	geneData := GeneData{
		UserID:        userID,
		EncryptedData: encryptedData,
		DataHash:      hashBytes,
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
