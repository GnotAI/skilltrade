package userskills

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

  "github.com/GnotAI/skilltrade/internal/users"
  "github.com/GnotAI/skilltrade/internal/skills"
)

type UserSkill struct {
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID  uuid.UUID `gorm:"type:uuid;not null"`
	SkillID uuid.UUID `gorm:"type:uuid;not null"`
	Type    string    `gorm:"type:text;not null;check:type IN ('offering', 'seeking')"`

	User  users.User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Skill skills.Skill `gorm:"foreignKey:SkillID;constraint:OnDelete:CASCADE"`

	// Ensuring uniqueness for (user_id, skill_id, type)
  _ struct{} `gorm:"uniqueIndex:user_skill_unique,unique"`
}

// BeforeCreate hook to generate UUID for new user_skills entries
func (us *UserSkill) BeforeCreate(tx *gorm.DB) (err error) {
	if us.ID == uuid.Nil {
		us.ID = uuid.New()
	}
	return
}
