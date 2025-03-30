package db

import (
  "gorm.io/gorm"
	"github.com/GnotAI/skilltrade/internal/sessions"
	"github.com/GnotAI/skilltrade/internal/trades"
	"github.com/GnotAI/skilltrade/internal/users"
)

// MigrateUsersTable ensures the users table is created in the database
func MigrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
    &users.User{},
    &sessions.Session{},
    &trades.TradeRequest{},
  )
}
