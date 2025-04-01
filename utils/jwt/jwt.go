package jwtutil

import (
  "os"
	"time"
  "errors"

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

// ParseToken takes a JWT token string, validates it and returns claims
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify that the token's signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	// Extract the claims
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}
