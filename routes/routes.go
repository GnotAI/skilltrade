package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/GnotAI/skilltrade/middleware"
	"github.com/GnotAI/skilltrade/utils/logger"
	"github.com/go-chi/chi/v5"
)



func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

  log := logger.NewLogger()

  // Initialize middleware
  r.Use(middleware.LoggerMiddleware(log))
	r.Use(middleware.RateLimiterMiddleware())

	// Mount user routes
  r.Mount("/auth", authRoutes())
  r.Mount("/skills", skillsRoutes(skillHandler))
	// r.Mount("/users", userRoutes(userService))
  r.Mount("/docs", docsRoute())

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
