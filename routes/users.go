package routes

import (
	"github.com/GnotAI/skilltrade/db"
	"github.com/GnotAI/skilltrade/internal/users"
	"github.com/GnotAI/skilltrade/internal/userskills"
	"github.com/GnotAI/skilltrade/middleware"
	"github.com/go-chi/chi/v5"
)

// Initialize Repositories
var userRepo = users.NewUserRepository(db.DB)
var userSkillRepo = userskills.NewUserSkillRepository(db.DB)

// Initialize services
var userService = users.NewUserService(userRepo)
var userSkillService = userskills.NewUserSkillService(userSkillRepo)

// Initialize handlers
var userHandler = users.NewUserHandler(userService)
var userSkillHandler = userskills.NewUserSkillHandler(userSkillService)

func userRoutes(userService *users.UserService) *chi.Mux {
	r := chi.NewRouter()

  r.Group(func(r chi.Router) {
    // Apply the AuthMiddleware for this group
    r.Use(middleware.AuthMiddleware)  

    // The refresh token endpoint now requires the user to be authenticated
    r.Post("/skills", userSkillHandler.CreateUserSkillHandler) 
  })
  
	r.Post("/", userHandler.CreateUserHandler)
	r.Put("/{id}", userHandler.UpdateUserHandler)
	r.Delete("/del/{id}", userHandler.DeleteUserHandler)

	return r
}
