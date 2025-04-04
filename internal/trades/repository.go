package trades

import (
	"context"
  "github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateTrade(ctx context.Context, trade *TradeRequest) error
	GetTradesByUserID(ctx context.Context, userID uuid.UUID) ([]TradeRequest, error)
  UpdateTradeStatus(ctx context.Context, tradeID uuid.UUID, status string) error
}

type repository struct {
	db *gorm.DB
}

func NewTradeRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateTrade(ctx context.Context, trade *TradeRequest) error {
	return r.db.WithContext(ctx).Create(trade).Error
}

func (r *repository) GetTradesByUserID(ctx context.Context, userID uuid.UUID) ([]TradeRequest, error) {
	var trades []TradeRequest
	err := r.db.WithContext(ctx).
		Where("sender_id = ? OR receiver_id = ?", userID, userID).
		Preload("Sender").
		Preload("Receiver").
		Preload("SenderSkill").
		Preload("ReceiverSkill").
		Find(&trades).Error
	return trades, err
}

func (r *repository) UpdateTradeStatus(ctx context.Context, tradeID uuid.UUID, status string) error {
	return r.db.WithContext(ctx).
		Model(&TradeRequest{}).
		Where("id = ?", tradeID).
		Update("status", status).Error
}
