package users

import (
	"time"
)

type User struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email     string    `gorm:"type:text;unique;not null"`
	Password  string    `gorm:"type:text;not null"`
	FullName  string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()"`
}
