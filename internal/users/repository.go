package users

import (
	"gorm.io/gorm"
)

// MigrateUsersTable ensures the users table is created in the database
func MigrateUsersTable(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
