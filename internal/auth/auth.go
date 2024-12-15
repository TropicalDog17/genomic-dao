package auth

import (
	"math/rand"
	"regexp"

	"gorm.io/gorm"
)

type authService struct {
	UserRepository UserRepository
}
type AuthService interface {
	Register(address string) (uint32, error)
	Authenticate(address string) (*User, error)
}

func NewAuthService(userRepository UserRepository) AuthService {
	return &authService{
		UserRepository: userRepository,
	}
}

func (s *authService) Register(address string) (uint32, error) {
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

func (s *authService) Authenticate(address string) (*User, error) {
	// sanitize the address
	if ok := ValidateAddress(address); !ok {
		return nil, ErrInvalidAddress
	}

	user, err := s.UserRepository.FindByPubkey(address)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

// ValidateAddress checks if the address is valid Ethereum address
func ValidateAddress(address string) bool {
	if len(address) == 0 {
		return false
	}

	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}
