package tee

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// TeeService provides minimal TEE functionality
type teeService struct{}

type TeeService interface {
	CalculateRiskScore(encryptedData []byte, key []byte) (int, error)
	ProcessGeneData(data []byte, publicKey []byte) ([]byte, error)
	DecryptAndVerify(encryptedData []byte, publicKey []byte) ([]byte, error)
}

// NewTeeService creates a new TEE instance
func NewTeeService() TeeService {
	return &teeService{}
}

// GeneticData represents processed genetic markers
type GeneticData struct {
	Markers []float64
	Hash    [32]byte
}

// CalculateRiskScore computes genetic risk score within TEE
func (t *teeService) CalculateRiskScore(encryptedData []byte, key []byte) (riskScore int, err error) {
	// Decrypt and verify data
	data, err := t.DecryptAndVerify(encryptedData, key)
	if err != nil {
		return 0, err
	}

	// Convert decrypted data to genetic markers
	markers, err := bytesToMarkers(data)
	if err != nil {
		return 0, err
	}
	var rawRiskScore float64
	var weights = []float64{0.1, 0.2, 0.3, 0.4}

	for i, marker := range markers {
		if i < len(weights) {
			rawRiskScore += marker * weights[i]
		}
	}

	// normalize risk score to be 1,2,3,4
	if rawRiskScore < 0.25 {
		riskScore = 1
	} else if rawRiskScore < 0.5 {
		riskScore = 2
	} else if rawRiskScore < 0.75 {
		riskScore = 3
	} else {
		riskScore = 4
	}

	return riskScore, nil
}

// ProcessGeneData processes and protects gene data within the TEE
func (t *teeService) ProcessGeneData(data []byte, publicKey []byte) ([]byte, error) {
	// Derive a proper 32-byte key using SHA-256
	key := sha256.Sum256(publicKey)

	// Calculate hash of original data
	hash := sha256.Sum256(data)

	// Create AES-256 cipher
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, fmt.Errorf("creating cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("creating GCM: %w", err)
	}

	// Generate random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("generating nonce: %w", err)
	}

	// Combine hash and data
	combined := append(hash[:], data...)

	// Encrypt the combined data
	ciphertext := gcm.Seal(nonce, nonce, combined, nil)
	return ciphertext, nil
}

// DecryptAndVerify decrypts data and verifies its integrity
func (t *teeService) DecryptAndVerify(encryptedData []byte, publicKey []byte) ([]byte, error) {
	// Input validation
	if len(encryptedData) == 0 {
		return nil, errors.New("empty encrypted data")
	}

	// Derive the same 32-byte key using SHA-256
	key := sha256.Sum256(publicKey)

	// Create AES-256 cipher
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, fmt.Errorf("creating cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("creating GCM: %w", err)
	}

	// Extract nonce
	nonceSize := gcm.NonceSize()
	if len(encryptedData) < nonceSize {
		return nil, errors.New("invalid ciphertext: too short")
	}

	nonce := encryptedData[:nonceSize]
	ciphertext := encryptedData[nonceSize:]

	// Decrypt the data
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}

	// Verify data format
	if len(plaintext) < sha256.Size {
		return nil, errors.New("invalid data format: missing hash")
	}

	// Extract hash and data
	storedHash := plaintext[:sha256.Size]
	data := plaintext[sha256.Size:]

	// Calculate and verify hash
	calculatedHash := sha256.Sum256(data)
	if !compareHashes(calculatedHash[:], storedHash) {
		return nil, errors.New("data integrity verification failed")
	}

	return data, nil
}

// bytesToMarkers converts byte slice to float64 markers
func bytesToMarkers(data []byte) ([]float64, error) {
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
