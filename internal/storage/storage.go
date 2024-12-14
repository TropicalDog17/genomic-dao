package storage

type GeneDataStorageService struct {
	geneDataRepository GenDataRepository
}

func NewGeneDataStorageService(repo GenDataRepository) *GeneDataStorageService {
	return &GeneDataStorageService{
		geneDataRepository: repo,
	}
}

func (s *GeneDataStorageService) StoreGeneData(userID uint64, encryptedData []byte, signatureBytes []byte, hashBytes []byte) (string, error) {
	return s.geneDataRepository.StoreGeneData(userID, encryptedData, signatureBytes, hashBytes)
}

func (s *GeneDataStorageService) RetrieveGeneData(fileID string) ([]byte, error) {
	return s.geneDataRepository.RetrieveGeneData(fileID)
}
