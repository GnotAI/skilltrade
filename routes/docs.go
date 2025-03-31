package routes

import (
  "os"
  "net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func DocsRoute() *chi.Mux {
  r := chi.NewRouter()
	serverURL := os.Getenv("SERVER_URL")
	if serverURL == "" {
		serverURL = "http://localhost:8080"
	}

	// Serve OpenAPI YAML file
	r.Get("/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "openapi.yaml")
	})

	// Swagger UI loads OpenAPI spec from the correct server URL
	r.Get("/*", httpSwagger.Handler(
		httpSwagger.URL(serverURL+"/openapi.yaml"), // Use dynamic server URL
	))

  return r
}
