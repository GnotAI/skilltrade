package middleware

import (
  "time"
	"context"
	"net/http"
	"strings"

	jwtutil "github.com/GnotAI/skilltrade/utils/jwt"
)

// AuthMiddleware is a middleware to check JWT validity
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
    authHeader = strings.TrimPrefix(authHeader, "Bearer ")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		tokenString := authHeader

		// Parse and validate the token using the JWT utility
		claims, err := jwtutil.ParseToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Validate the token's expiration (and other claims if needed)
		if claims.ExpiresAt.Before(time.Now()) {
			http.Error(w, "Token has expired", http.StatusUnauthorized)
			return
		}

		// Store the user information in the request context for later use
		// This can be useful to access the authenticated user in the handler
    ctx := context.WithValue(r.Context(), "AuthorizationToken", tokenString) 
		r = r.WithContext(ctx)

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
