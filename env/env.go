package env

import (
  "log"

  gde "github.com/joho/godotenv"
)

func init() {
  if err := gde.Load(); err != nil {
    log.Fatalf("Error loading env file: %v", err)
  }
}
