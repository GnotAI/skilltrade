package email

import (
	"errors"
	"regexp"
)

// ValidateEmail checks if the provided email matches the SQL regex constraint
func ValidateEmail(email string) error {
	// Case-insensitive regex to match the SQL pattern
	emailRegex := `(?i)^.+@.+\..+$`
	matched, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("invalid email format")
	}
	return nil
}
