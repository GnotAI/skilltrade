package docs

import (
  "os"
  "net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

var SERV_URL = os.Getenv("SERVER_URL")

func Serve(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/swagger.yaml")
	}

var SwaggerHandler = httpSwagger.Handler(
		httpSwagger.URL(SERV_URL + "/swagger.yaml"), // The URL to the generated Swagger JSON file
)
