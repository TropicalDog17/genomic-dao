package auth

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID     uint32 `gorm:"primaryKey"`
	Pubkey string `gorm:"unique,not null"`
}

type UserRepository interface {
	Validate(user *User) error
	Create(user *User) error
	FindByPubkey(Pubkey string) (*User, error)
	FindByUserID(userID uint32) (*User, error)
}

// Implement the UserRepository interface
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func validateUser(user *User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	if len(user.Pubkey) == 0 {
		return errors.New("pubkey is empty")
	}
	return nil
}

func (r *userRepository) Validate(user *User) error {
	return validateUser(user)
}

func (r *userRepository) Create(user *User) error {
	if err := r.Validate(user); err != nil {
		return err
	}
	return r.db.Create(user).Error
}

func (r *userRepository) FindByPubkey(p string) (*User, error) {
	var user User
	if err := r.db.Where("pubkey = ?", p).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUserID(userID uint32) (*User, error) {
	var user User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
