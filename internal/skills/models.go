package skills

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Skill struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"type:text;not null"`
}

// BeforeCreate hook to generate UUID for new users
func (sk *Skill) BeforeCreate(tx *gorm.DB) (err error) {
	if sk.ID == uuid.Nil {
		sk.ID = uuid.New()
	}
	return
}
