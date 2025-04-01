package auth

import (
	"github.com/GnotAI/skilltrade/internal/users"
)

type AuthRepository struct {
  UserRepo *users.UserRepository
}

func NewAuthRepository(userRepo *users.UserRepository) *AuthRepository {
	return &AuthRepository{UserRepo: userRepo}
}

// CreateUser saves a new user in the database by calling UserRepository
func (r *AuthRepository) CreateUser(user *users.User) error {
	return r.UserRepo.CreateUser(user)
}

// GetUserByEmail retrieves a user by email
func (r *AuthRepository) GetUserByEmail(email string) (*users.User, error) {
	return r.UserRepo.GetUserByEmail(email)
}
