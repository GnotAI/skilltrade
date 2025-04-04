package main

import (
  db "github.com/GnotAI/skilltrade/db"
	"github.com/GnotAI/skilltrade/routes"
)

func main() {
  defer db.DisconnectDB()

  // StartServer
  routes.StartServer()
}
