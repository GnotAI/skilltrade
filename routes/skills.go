package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/GnotAI/skilltrade/internal/skills"
	"github.com/GnotAI/skilltrade/db"
)

// Initialize Repository
var skillRepo = skills.NewSkillRepository(db.DB)

// Initialize Service
var skillService = skills.NewSkillService(skillRepo)

// Initialize Handler
var skillHandler = skills.NewSkillHandler(skillService)


func skillsRoutes(authService *skills.SkillService) *chi.Mux {
	r := chi.NewRouter()

  r.Get("/", skillHandler.GetAllSkillsHandler) 
	// r.Post("/signup", authHandler.SignUpHandler)
	// r.Post("/signin", authHandler.SignInHandler)

	return r
}
