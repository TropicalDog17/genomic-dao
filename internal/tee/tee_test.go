package tee

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/binary"
	"testing"
)

func TestTeeService_ProcessAndDecrypt(t *testing.T) {
	// Setup
	service := NewTeeService()

	// Generate key pair
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	// Test cases
	tests := []struct {
		name        string
		data        []byte
		expectError bool
		riskLevel   int
	}{
		{
			name:        "Valid data low risk",
			data:        generateTestData(5, 0.2),
			expectError: false,
			riskLevel:   1,
		},
		{
			name:        "Valid data medium risk",
			data:        generateTestData(5, 0.4),
			expectError: false,
			riskLevel:   2,
		},
		{
			name:        "Valid data high risk",
			data:        generateTestData(5, 0.8),
			expectError: false,
			riskLevel:   4,
		},
		{
			name:        "Empty data",
			data:        []byte{},
			expectError: true,
		},
		{
			name:        "Invalid size data",
			data:        []byte{1, 2, 3}, // Not multiple of 8
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encrypt
			encrypted, riskLevel, err := service.ProcessAndEncrypt(tt.data, &privKey.PublicKey)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error during encryption: %v", err)
				return
			}

			if riskLevel != tt.riskLevel {
				t.Errorf("Expected risk level %d, got %d", tt.riskLevel, riskLevel)
			}

			// Decrypt
			decrypted, err := service.DecryptData(encrypted, privKey)
			if err != nil {
				t.Errorf("Unexpected error during decryption: %v", err)
				return
			}

			// Verify decrypted data matches original
			if len(decrypted) != len(tt.data) {
				t.Errorf("Decrypted data length mismatch. Expected %d, got %d", len(tt.data), len(decrypted))
				return
			}

			for i := range tt.data {
				if decrypted[i] != tt.data[i] {
					t.Errorf("Decrypted data mismatch at position %d. Expected %d, got %d", i, tt.data[i], decrypted[i])
					return
				}
			}
		})
	}
}

func TestTeeService_DecryptData_InvalidInput(t *testing.T) {
	service := NewTeeService()
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	tests := []struct {
		name        string
		data        []byte
		expectError bool
	}{
		{
			name:        "Empty data",
			data:        []byte{},
			expectError: true,
		},
		{
			name:        "Too short data",
			data:        make([]byte, 31), // Less than hash size
			expectError: true,
		},
		{
			name:        "Invalid encrypted data",
			data:        make([]byte, 100), // Random invalid data
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.DecryptData(tt.data, privKey)
			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

// Helper function to generate test data with specific risk level
func generateTestData(count int, riskFactor float64) []byte {
	data := make([]byte, count*8)
	value := uint64(riskFactor * float64(1<<63))

	for i := 0; i < count; i++ {
		binary.LittleEndian.PutUint64(data[i*8:], value)
	}

	return data
}
