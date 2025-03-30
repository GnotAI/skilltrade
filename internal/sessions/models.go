package sessions

import (
	"time"

  "gorm.io/gorm"
	"github.com/google/uuid"
  "github.com/GnotAI/skilltrade/internal/trades"
)

type Session struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TradeID     uuid.UUID `gorm:"type:uuid;not null"`
	ScheduledAt time.Time `gorm:"not null"`
	Completed   bool      `gorm:"default:false"`

	Trade trades.TradeRequest `gorm:"foreignKey:TradeID;constraint:OnDelete:CASCADE"`
}

// BeforeCreate hook to generate UUID
func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return
}
