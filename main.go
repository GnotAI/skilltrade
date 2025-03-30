package main

import (
	_ "github.com/GnotAI/skilltrade/env"
  db "github.com/GnotAI/skilltrade/db"
	"github.com/GnotAI/skilltrade/routes"
)

func main() {
  defer db.DisconnectDB()

  // StartServer
  routes.StartServer()
}
