package hash

import (
  "errors"
  "regexp"
	"golang.org/x/crypto/bcrypt"
)

// ValidatePassword checks if a password meets security requirements
func ValidatePassword(password string) error {
	// At least 8 characters
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Must contain at least one special character
	specialCharPattern := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`)
	if !specialCharPattern.MatchString(password) {
		return errors.New("password must contain at least one special character")
	}

	// Must contain at least one number
	numberPattern := regexp.MustCompile(`[0-9]`)
	if !numberPattern.MatchString(password) {
		return errors.New("password must contain at least one number")
	}

	// Must contain at least one uppercase letter
	uppercasePattern := regexp.MustCompile(`[A-Z]`)
	if !uppercasePattern.MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	return nil
}

// HashPassword hashes a plaintext password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// ComparePassword compares a hashed password with a plaintext password.
func ComparePassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
