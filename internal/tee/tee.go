package tee

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

type TeeService interface {
	ProcessAndEncrypt(data []byte, pubKey *ecdsa.PublicKey) ([]byte, int, error)
	DecryptData(encryptedData []byte, privKey *ecdsa.PrivateKey) ([]byte, error)
}

type teeService struct{}

func NewTeeService() TeeService {
	return &teeService{}
}

// deriveKey creates a consistent key from either public or private key
func deriveKey(key interface{}) []byte {
	switch k := key.(type) {
	case *ecdsa.PrivateKey:
		combined := append(k.X.Bytes(), k.Y.Bytes()...)
		combined = append(combined, k.D.Bytes()...)
		h := sha256.Sum256(combined)
		return h[:]
	case *ecdsa.PublicKey:
		combined := append(k.X.Bytes(), k.Y.Bytes()...)
		h := sha256.Sum256(combined)
		return h[:]
	default:
		return nil
	}
}

func (t *teeService) DecryptData(encryptedData []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	if len(encryptedData) < sha256.Size+12 {
		return nil, fmt.Errorf("encrypted data too short")
	}

	hash := encryptedData[:sha256.Size]
	nonce := encryptedData[sha256.Size : sha256.Size+12]
	ciphertext := encryptedData[sha256.Size+12:]

	key := deriveKey(&privKey.PublicKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("creating cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("creating GCM: %w", err)
	}

	decryptedData, err := aesgcm.Open(nil, nonce, ciphertext, hash)
	if err != nil {
		return nil, fmt.Errorf("decrypting: %w", err)
	}

	actualHash := sha256.Sum256(decryptedData)
	if !compareHashes(hash[:], actualHash[:]) {
		return nil, fmt.Errorf("data integrity check failed")
	}

	return decryptedData, nil
}

func (t *teeService) ProcessAndEncrypt(data []byte, pubKey *ecdsa.PublicKey) ([]byte, int, error) {
	markers, err := bytesToMarkers(data)
	if err != nil {
		return nil, 0, fmt.Errorf("invalid marker data: %w", err)
	}

	if len(markers) == 0 {
		return nil, 0, fmt.Errorf("no markers provided")
	}

	var sum float64
	for _, marker := range markers {
		sum += marker
	}
	score := sum / float64(len(markers))

	riskLevel := 1
	switch {
	case score < 0.25:
		riskLevel = 1
	case score < 0.50:
		riskLevel = 2
	case score < 0.75:
		riskLevel = 3
	default:
		riskLevel = 4
	}

	hash := sha256.Sum256(data)

	// Use consistent key derivation
	key := deriveKey(pubKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, 0, fmt.Errorf("creating cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, 0, fmt.Errorf("creating GCM: %w", err)
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, 0, fmt.Errorf("generating nonce: %w", err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, data, hash[:])

	result := make([]byte, sha256.Size+len(nonce)+len(ciphertext))
	copy(result[:sha256.Size], hash[:])
	copy(result[sha256.Size:sha256.Size+len(nonce)], nonce)
	copy(result[sha256.Size+len(nonce):], ciphertext)

	return result, riskLevel, nil
}

func bytesToMarkers(data []byte) ([]float64, error) {
	markers := make([]float64, len(data)/8)
	for i := range markers {
		markers[i] = float64(binary.LittleEndian.Uint64(data[i*8:(i+1)*8])) / float64(1<<63)
	}
	return markers, nil
}

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
