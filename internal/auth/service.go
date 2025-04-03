package auth

import (
	"errors"

	"github.com/GnotAI/skilltrade/internal/users"
	"github.com/GnotAI/skilltrade/utils/hash"
	"github.com/GnotAI/skilltrade/utils/email"
	"github.com/GnotAI/skilltrade/utils/jwt"
)

type AuthService struct {
	Repo *AuthRepository
}

func NewAuthService(repo *AuthRepository) *AuthService {
	return &AuthService{Repo: repo}
}

// SignUp registers a new user
func (s *AuthService) SignUp(user *users.User) error {

	// Validate email format before proceeding
	if err := email.ValidateEmail(user.Email); err != nil {
		return err 
	}

  if user.Email == "" || user.Password == "" || user.FullName == "" {
    return errors.New("Missing email, password or full name")
  }

	// Check if user already exists
	existingUser, err := s.Repo.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("user already exists")
	}

  // Validate password before hashing
  if err := hash.ValidatePassword(user.Password); err != nil {
    return err 
  }

	// Hash password
	hashedPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	err = s.Repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

// SignIn authenticates a user and returns a JWT token
func (s *AuthService) SignIn(email, password string) (string, error) {

  if email == "" || password == "" {
    return "", errors.New("Missing email or password")
  }

	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("User with email doesn't exist")
	}

	// Check password
	if !hash.ComparePassword(user.Password, password) {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := jwtutil.GenerateToken(user.ID.String())
	if err != nil {
		return "", err
	}

	return token, nil
}

// RefreshToken generates a new JWT token if the current one is valid
func (s *AuthService) RefreshToken(oldToken string) (string, error) {
	claims, err := jwtutil.ParseToken(oldToken)
	if err != nil {
		return "", errors.New("invalid token")
	}

	// Generate new token
	newToken, err := jwtutil.GenerateToken(claims.UserID)
	if err != nil {
		return "", err
	}

	return newToken, nil
}
