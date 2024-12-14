package tee

import (
	"crypto/rand"
	"testing"
)

func TestNewTeeService(t *testing.T) {
	service := NewTeeService()
	if service == nil {
		t.Error("NewTeeService returned nil")
	}
}

func TestProcessGeneData(t *testing.T) {
	service := NewTeeService()

	// Generate test key (32 bytes for AES-256)
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		t.Fatalf("Failed to generate test key: %v", err)
	}

	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{
			name:    "Valid data",
			data:    []byte("test genetic data"),
			wantErr: false,
		},
		{
			name:    "Empty data",
			data:    []byte{},
			wantErr: false,
		},
		{
			name:    "Nil data",
			data:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encrypted, err := service.ProcessGeneData(tt.data, key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessGeneData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && encrypted == nil {
				t.Error("ProcessGeneData() returned nil encrypted data")
			}
		})
	}
}
