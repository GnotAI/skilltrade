package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/GnotAI/skilltrade/internal/sessions"
	"github.com/GnotAI/skilltrade/db"
	"github.com/GnotAI/skilltrade/middleware"
)

// Initialize Repository
var sessRepo = sessions.NewRepository(db.DB)

// Initialize Service
var sessService = sessions.NewService(sessRepo, tradeRepo)

// Initialize Handler
var sessHandler = sessions.NewSessionHandler(sessService)

func sessRoutes() *chi.Mux {
	r := chi.NewRouter()

  r.Group(func(r chi.Router) {
    r.Use(middleware.AuthMiddleware)  

    r.Post("/", sessHandler.ScheduleSession) 
    r.Patch("/{id}/complete", sessHandler.MarkSessionCompleted) 
  })

	return r
}
