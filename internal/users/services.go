package users

import (
	"fmt"

	"github.com/google/uuid"
)

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// CreateUser handles user creation
func (s *UserService) CreateUser(email, password, fullName string) error {
	user := &User{
		Email:     email,
		Password:  password, // Hashing should be done before calling Repo.CreateUser
		FullName:  fullName,
	}
	return s.Repo.CreateUser(user)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(userID uuid.UUID, email, password, fullName string) error {
  _, err := s.Repo.GetUserByID(userID)
  if err != nil {
    return fmt.Errorf("user not found")
  }

	req := &User{
		ID:        userID,
		Email:     email,
		Password:  password, // Hashing should be done before calling Repo.CreateUser
		FullName:  fullName,
	}

	return s.Repo.UpdateUser(req)
}

// DeleteUser removes a user by ID
func (s *UserService) DeleteUser(userID uuid.UUID) error {
  _, err := s.Repo.GetUserByID(userID)
  if err != nil {
    return fmt.Errorf("user not found")
  }

  return s.Repo.DeleteUser(userID)
}
