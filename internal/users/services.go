package users

import (
  "fmt"

  "github.com/GnotAI/skilltrade/utils/hash"
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
  // Validate password before hashing
  if err := hash.ValidatePassword(password); err != nil {
    return err 
  }

  // Hash the password before storing
  hashedPassword, err := hash.HashPassword(password)
  if err != nil {
    return err
  }

  user := &User{
    Email:     email,
    Password:  hashedPassword, // Hashing should be done before calling Repo.CreateUser
    FullName:  fullName,
  }
  return s.Repo.CreateUser(user)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(userID uuid.UUID, email, password, fullName string) error {
  existingUser, err := s.Repo.GetUserByID(userID)
  if err != nil {
    return fmt.Errorf("user not found")
  }

  req := &User{
    ID:        userID,
    Email: existingUser.Email,
    Password: existingUser.Password,
    FullName: existingUser.FullName,
    CreatedAt: existingUser.CreatedAt,
  }

  // Update only the fields that are provided
  if email != "" {
    req.Email = email
	}

	if fullName != "" {
    req.FullName = fullName
	}

	if password != "" {
    if err := hash.ValidatePassword(password); err != nil {
      return err
    }

		hashedPassword, err := hash.HashPassword(password)
		if err != nil {
			return err
		}
		req.Password = hashedPassword
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
