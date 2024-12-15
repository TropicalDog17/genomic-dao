package auth

import (
	"math/rand"
	"regexp"

	"gorm.io/gorm"
)

type AuthService struct {
	UserRepository UserRepository
}

func NewAuthService(userRepository UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (s *AuthService) Register(address string) (uint32, error) {
	// sanitize the address
	if ok := ValidateAddress(address); !ok {
		return 0, ErrInvalidAddress
	}
	userID := rand.Uint32()
	user := &User{ID: userID, Pubkey: address}
	if err := s.UserRepository.Validate(user); err != nil {
		return 0, err
	}
	if _, err := s.UserRepository.FindByPubkey(address); err == nil {
		return 0, ErrUserExists
	}
	if err := s.UserRepository.Create(user); err != nil {
		return 0, err
	}
	return userID, nil
}

func (s *AuthService) Authenticate(userID uint32, address string) error {
	// sanitize the address
	if ok := ValidateAddress(address); !ok {
		return ErrInvalidAddress
	}

	user, err := s.UserRepository.FindByPubkey(address)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrUserNotFound
		}
		return err
	}

	if user.ID != userID {
		return ErrUnauthorized
	}
	return nil
}

// ValidateAddress checks if the address is valid Ethereum address
func ValidateAddress(address string) bool {
	if len(address) == 0 {
		return false
	}

	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}
