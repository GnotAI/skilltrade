package trades

import (
	"time"

  "gorm.io/gorm"
	"github.com/google/uuid"
  "github.com/GnotAI/skilltrade/internal/users"
  "github.com/GnotAI/skilltrade/internal/skills"
)

type TradeRequest struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	SenderID        uuid.UUID `gorm:"type:uuid;not null"`
	ReceiverID      uuid.UUID `gorm:"type:uuid;not null"`
	SenderSkillID   uuid.UUID `gorm:"type:uuid;not null"`
	ReceiverSkillID uuid.UUID `gorm:"type:uuid;not null"`
	Status          string    `gorm:"type:text;default:'pending';check:status IN ('pending', 'accepted', 'rejected')"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"`

	Sender        users.User  `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE"`
	Receiver      users.User  `gorm:"foreignKey:ReceiverID;constraint:OnDelete:CASCADE"`
	SenderSkill   skills.Skill `gorm:"foreignKey:SenderSkillID;constraint:OnDelete:CASCADE"`
	ReceiverSkill skills.Skill `gorm:"foreignKey:ReceiverSkillID;constraint:OnDelete:CASCADE"`
}

// BeforeCreate hook to generate UUID
func (t *TradeRequest) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return
}
