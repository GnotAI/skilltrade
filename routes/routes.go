package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/GnotAI/skilltrade/db"
	"github.com/GnotAI/skilltrade/internal/users"
	"github.com/go-chi/chi/v5"
)

// Initialize Repositories
var userRepo = users.NewUserRepository(db.DB)

// Initialize services
var userService = users.NewUserService(userRepo)

// Initialize handlers
var userHandler = users.NewUserHandler(userService)

func UserRoutes(userService *users.UserService) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", userHandler.CreateUserHandler)
	r.Put("/{id}", userHandler.UpdateUserHandler)
	r.Delete("/{id}", userHandler.DeleteUserHandler)

	return r
}

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()


	// Mount user routes
	r.Mount("/users", UserRoutes(userService))

	return r
}

func StartServer() {
  port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if PORT is not set
	}

	r := RegisterRoutes()

  srv := http.Server{
    Addr: ":" + port,
    Handler: r,
  }

	log.Printf("Server starting on port %s...", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
