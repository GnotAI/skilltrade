package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/GnotAI/skilltrade/db"
	"github.com/GnotAI/skilltrade/internal/users"
)

// Initialize Repositories
var UserRepo = users.NewUserRepository(db.DB)

// Initialize services
var userService = users.NewUserService(UserRepo)

// Initialize handlers
var userHandler = users.NewUserHandler(userService)

func UserRoutes(userService *users.UserService) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", userHandler.CreateUserHandler)
	r.Put("/{id}", userHandler.UpdateUserHandler)
	r.Delete("/del/{id}", userHandler.DeleteUserHandler)

	return r
}
