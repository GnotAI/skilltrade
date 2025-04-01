package jwtutil

import (
  "os"
	"time"
  "errors"
  "strings"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key for signing JWTs (change this in production)
var JWT_SEC = os.Getenv("JWT_SEC")
var jwtSecret = []byte(JWT_SEC)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT token for a given user ID
func GenerateToken(userID string) (string, error) {
	// Set custom claims with the user ID
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "skilltrade",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)), // token expires in 15 minutes
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token
	return token.SignedString(jwtSecret)
}

// ParseToken verifies and extracts claims from a JWT token
func ParseToken(tokenString string) (*Claims, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Check if the token string is empty after trimming
	if tokenString == "" {
		return nil, errors.New("token is required")
	}

	// Parse the token using your JWT secret key
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Return the secret key used for signing the token
		return jwtSecret, nil
	})

	// Check for errors during token parsing
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, errors.New("could not parse token")
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// If everything is fine, return the claims
	return claims, nil
}
