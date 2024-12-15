package auth

import (
	"bytes"
	"errors"
	"testing"

	"gorm.io/gorm"
)

var (
	validAddress   = "0x1234567890123456789012345678901234567890"
	invalidAddress = "0xinvalid"
)

// Mock repository for testing
type mockUserRepository struct {
	users map[uint32]*User
	err   error
}

func newMockUserRepository() *mockUserRepository {
	return &mockUserRepository{
		users: make(map[uint32]*User),
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

func (m *mockUserRepository) FindByPubkey(pubkey string) (*User, error) {
	if m.err != nil {
		return nil, m.err
	}
	for _, user := range m.users {
		if bytes.Equal([]byte(pubkey), []byte(user.Pubkey)) {
			return user, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockUserRepository) FindByUserID(userID uint32) (*User, error) {
	if m.err != nil {
		return nil, m.err
	}
	if user, exists := m.users[userID]; exists {
		return user, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func TestAuthService_Register(t *testing.T) {
	tests := []struct {
		name    string
		address string
		repoErr error
		wantErr error
	}{
		{
			name:    "successful registration",
			address: validAddress,
			repoErr: nil,
			wantErr: nil,
		},
		{
			name:    "invalid address",
			address: invalidAddress,
			repoErr: nil,
			wantErr: ErrInvalidAddress,
		},
		{
			name:    "empty address",
			address: "",
			repoErr: nil,
			wantErr: ErrInvalidAddress,
		},
		{
			name:    "repository error",
			address: validAddress,
			repoErr: errors.New("db error"),
			wantErr: errors.New("db error"),
		},
		{
			name:    "user already exists",
			address: validAddress,
			repoErr: nil,
			wantErr: ErrUserExists,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := newMockUserRepository()
			mockRepo.err = tt.repoErr

			// For "user already exists" test
			if tt.name == "user already exists" {
				mockRepo.users[1] = &User{
					ID:     1,
					Pubkey: tt.address,
				}
			}

			service := NewAuthService(mockRepo)

			userID, err := service.Register(tt.address)

			// Check error
			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check userID for successful cases
			if tt.wantErr == nil {
				if userID == 0 {
					t.Error("Register() returned userID = 0, want non-zero")
				}

				// Verify user was created in repository
				user, err := mockRepo.FindByPubkey(tt.address)
				if err != nil {
					t.Errorf("Failed to find created user: %v", err)
				}
				if user.ID != userID {
					t.Errorf("Created user ID = %v, want %v", user.ID, userID)
				}
			}
		})
	}
}

func TestAuthService_Authenticate(t *testing.T) {
	tests := []struct {
		name      string
		setupRepo func(*mockUserRepository)
		userID    uint32
		address   string
		wantErr   error
	}{
		{
			name: "successful authentication",
			setupRepo: func(m *mockUserRepository) {
				m.users[1] = &User{ID: 1, Pubkey: validAddress}
			},
			userID:  1,
			address: validAddress,
			wantErr: nil,
		},
		{
			name:      "invalid address format",
			setupRepo: func(m *mockUserRepository) {},
			userID:    1,
			address:   invalidAddress,
			wantErr:   ErrInvalidAddress,
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
				m.users[1] = &User{ID: 1, Pubkey: validAddress}
			},
			userID:  2,
			address: validAddress,
			wantErr: ErrUnauthorized,
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

			if (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr == nil) {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
