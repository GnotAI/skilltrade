package main

import (
  _ "github.com/GnotAI/skilltrade/env"
  db "github.com/GnotAI/skilltrade/db"
)

func main() {
  defer db.DisconnectDB()
}
