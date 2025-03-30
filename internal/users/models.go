package users

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email     string    `gorm:"type:text;unique;not null;check:email ~* '^.+@.+\\..+$'"`
	Password  string    `gorm:"type:text;not null"`
	FullName  string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()"`
}

// BeforeCreate hook to generate UUID for new users
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
