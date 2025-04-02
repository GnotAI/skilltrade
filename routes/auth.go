package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/GnotAI/skilltrade/internal/auth"
	"github.com/GnotAI/skilltrade/middleware"
)

// Initialize Repository
var authRepo = auth.NewAuthRepository(userRepo)

// Initialize Service
var authService = auth.NewAuthService(authRepo)

// Initialize Handler
var authHandler = auth.NewAuthHandler(authService)

func authRoutes(authService *auth.AuthService) *chi.Mux {
	r := chi.NewRouter()

  r.Group(func(r chi.Router) {
    // Apply the AuthMiddleware for this group
    r.Use(middleware.AuthMiddleware)  

    // The refresh token endpoint now requires the user to be authenticated
    r.Post("/refresh", authHandler.RefreshHandler) 
  })

	r.Post("/signup", authHandler.SignUpHandler)
	r.Post("/signin", authHandler.SignInHandler)

	return r
}
