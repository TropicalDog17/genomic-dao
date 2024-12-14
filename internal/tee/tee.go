package tee

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"io"
)

// TeeService provides minimal TEE functionality
type TeeService struct{}

// NewTeeService creates a new TEE instance
func NewTeeService() *TeeService {
	return &TeeService{}
}

// GeneticData represents processed genetic markers
type GeneticData struct {
	Markers []float64
	Hash    [32]byte
}

// CalculateRiskScore computes genetic risk score within TEE
func (t *TeeService) CalculateRiskScore(encryptedData []byte, key []byte) (float64, error) {
	// Decrypt and verify data
	data, err := t.decryptAndVerify(encryptedData, key)
	if err != nil {
		return 0, err
	}

	// Convert decrypted data to genetic markers
	markers, err := bytesToMarkers(data)
	if err != nil {
		return 0, err
	}
	var riskScore float64
	var weights = []float64{0.1, 0.2, 0.3, 0.4}

	for i, marker := range markers {
		if i < len(weights) {
			riskScore += marker * weights[i]
		}
	}

	// Normalize score between 0 and 1
	riskScore = riskScore / float64(len(markers))
	if riskScore < 0 {
		riskScore = 0
	}
	if riskScore > 1 {
		riskScore = 1
	}

	return riskScore, nil
}

// ProcessGeneData processes and protects gene data within the TEE
func (t *TeeService) ProcessGeneData(data []byte, key []byte) ([]byte, error) {
	// First calculate hash
	hash := sha256.Sum256(data)

	// Encrypt data using AES-GCM
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Combine hash and data before encryption
	combined := append(hash[:], data...)

	// Encrypt the combined data
	ciphertext := gcm.Seal(nonce, nonce, combined, nil)
	return ciphertext, nil
}

// decryptAndVerify decrypts data and verifies its integrity
func (t *TeeService) decryptAndVerify(encryptedData []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedData) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	// Separate hash and data
	if len(plaintext) < sha256.Size {
		return nil, errors.New("invalid data format")
	}
	storedHash := plaintext[:sha256.Size]
	data := plaintext[sha256.Size:]

	// Verify hash
	calculatedHash := sha256.Sum256(data)
	if !compareHashes(calculatedHash[:], storedHash) {
		return nil, errors.New("data integrity check failed")
	}

	return data, nil
}

// bytesToMarkers converts byte slice to float64 markers
func bytesToMarkers(data []byte) ([]float64, error) {
	if len(data)%8 != 0 {
		return nil, errors.New("invalid data length for markers")
	}

	markers := make([]float64, len(data)/8)
	for i := range markers {
		markers[i] = float64(binary.LittleEndian.Uint64(data[i*8:(i+1)*8])) / float64(1<<63)
	}
	return markers, nil
}

// compareHashes performs constant-time comparison of hashes
func compareHashes(h1, h2 []byte) bool {
	if len(h1) != len(h2) {
		return false
	}
	var diff byte
	for i := 0; i < len(h1); i++ {
		diff |= h1[i] ^ h2[i]
	}
	return diff == 0
}
