package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/GnotAI/skilltrade/internal/auth"
)

// Initialize Repository
var authRepo = auth.NewAuthRepository(UserRepo)

// Initialize Service
var authService = auth.NewAuthService(authRepo)

// Initialize Handler
var authHandler = auth.NewAuthHandler(authService)

func AuthRoutes(authService *auth.AuthService) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/signup", authHandler.SignUpHandler)
	r.Post("/signin", authHandler.SignInHandler)
	r.Post("/refresh", authHandler.RefreshHandler)

	return r
}
