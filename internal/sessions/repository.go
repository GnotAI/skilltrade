package sessions

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type sessRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *sessRepository {
	return &sessRepository{DB: db}
}

func (r *sessRepository) CreateSession(session *Session) error {
	return r.DB.Create(session).Error
}

func (r *sessRepository) GetSessionByID(id uuid.UUID) (*Session, error) {
	var session Session
	if err := r.DB.First(&session, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *sessRepository) GetSessionsByTradeID(tradeID uuid.UUID) ([]Session, error) {
	var sessions []Session
	if err := r.DB.Where("trade_id = ?", tradeID).Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *sessRepository) MarkSessionCompleted(sessionID uuid.UUID) error {
	return r.DB.Model(&Session{}).
		Where("id = ?", sessionID).
		Update("completed", true).Error
}
