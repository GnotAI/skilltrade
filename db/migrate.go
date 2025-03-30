package db

import (
	"github.com/GnotAI/skilltrade/internal/sessions"
	"github.com/GnotAI/skilltrade/internal/skills"
	"github.com/GnotAI/skilltrade/internal/trades"
	"github.com/GnotAI/skilltrade/internal/users"
	"github.com/GnotAI/skilltrade/internal/userskills"
	"gorm.io/gorm"
)

// MigrateUsersTable ensures the users table is created in the database
func MigrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
    &users.User{},
    &sessions.Session{},
    &trades.TradeRequest{},
    &skills.Skill{},
    &userskills.UserSkill{},
  )
}
