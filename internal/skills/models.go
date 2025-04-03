package skills

import (
	"github.com/google/uuid"
)

type Skill struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"type:text;not null"`
}
