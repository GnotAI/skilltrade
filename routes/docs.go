package routes

import (
	"github.com/GnotAI/skilltrade/docs"
	"github.com/go-chi/chi/v5"
)

func DocsRoute() *chi.Mux {
  r := chi.NewRouter()


	// Serve the raw swagger.yaml file
	r.Get("/swagger.yaml", docs.Serve)

	// Swagger UI route
	r.Get("/swagger/*", docs.SwaggerHandler)

  return r
}
