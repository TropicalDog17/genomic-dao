package auth

import (
	"bytes"
	"errors"
	"testing"

	"gorm.io/gorm"
)

var (
	validAddress = "0x1234567890123456789012345678901234567890"
)

// Mock repository for testing
type mockUserRepository struct {
	users map[uint64]*User
	err   error
}

func newMockUserRepository() *mockUserRepository {
	return &mockUserRepository{
		users: make(map[uint64]*User),
	}
}

func (m *mockUserRepository) Validate(user *User) error {
	return validateUser(user)
}

func (m *mockUserRepository) Create(user *User) error {
	if m.err != nil {
		return m.err
	}

	if err := m.Validate(user); err != nil {
		return err
	}

	m.users[user.ID] = user
	return nil
}

func (m *mockUserRepository) FindByPubkey(pubkey []byte) (*User, error) {
	if m.err != nil {
		return nil, m.err
	}
	for _, user := range m.users {
		if bytes.Equal(user.Pubkey, pubkey) {
			return user, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockUserRepository) FindByUserID(userID uint64) (*User, error) {
	if m.err != nil {
		return nil, m.err
	}
	if user, exists := m.users[userID]; exists {
		return user, nil
	}
	return nil, gorm.ErrRecordNotFound
}

// AuthService tests
func TestAuthService_Register(t *testing.T) {
	tests := []struct {
		name    string
		pubkey  []byte
		repoErr error
		wantErr bool
	}{
		{
			name:    "successful registration",
			pubkey:  []byte("valid-pubkey"),
			repoErr: nil,
			wantErr: false,
		},
		{
			name:    "empty pubkey",
			pubkey:  []byte{},
			repoErr: nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := newMockUserRepository()
			mockRepo.err = tt.repoErr
			service := NewAuthService(mockRepo)

			userID, err := service.Register(tt.pubkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && userID == 0 {
				t.Error("Register() returned userID = 0, want non-zero")
			}
		})
	}
}

func TestAuthService_Authenticate(t *testing.T) {
	tests := []struct {
		name      string
		setupRepo func(*mockUserRepository)
		userID    uint64
		address   string
		wantErr   error
	}{
		{
			name: "successful authentication",
			setupRepo: func(m *mockUserRepository) {
				m.users[1] = &User{ID: 1, Pubkey: []byte(validAddress)}
			},
			userID:  1,
			address: validAddress,
			wantErr: nil,
		},
		{
			name: "user not found",
			setupRepo: func(m *mockUserRepository) {
				// empty repo
			},
			userID:  1,
			address: validAddress,
			wantErr: ErrUserNotFound,
		},
		{
			name: "unauthorized - ID mismatch",
			setupRepo: func(m *mockUserRepository) {
				m.users[1] = &User{ID: 1, Pubkey: []byte("valid-address")}
			},
			userID:  2,
			address: validAddress,
			wantErr: ErrUserNotFound,
		},
		{
			name:      "empty address",
			setupRepo: func(m *mockUserRepository) {},
			userID:    1,
			address:   "",
			wantErr:   ErrInvalidAddress,
		},
		{
			name: "repository error",
			setupRepo: func(m *mockUserRepository) {
				m.err = errors.New("db error")
			},
			userID:  1,
			address: validAddress,
			wantErr: errors.New("db error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := newMockUserRepository()
			tt.setupRepo(mockRepo)
			service := NewAuthService(mockRepo)

			err := service.Authenticate(tt.userID, tt.address)
			if err == nil && tt.wantErr != nil {
				t.Errorf("Authenticate() error = nil, wantErr %v", tt.wantErr)
				return
			}
			if err != nil && tt.wantErr == nil {
				t.Errorf("Authenticate() error = %v, wantErr nil", err)
				return
			}
			if err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
