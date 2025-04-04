package trades

import (
	"context"
  "github.com/google/uuid"
	"gorm.io/gorm"
  "github.com/GnotAI/skilltrade/internal/userskills"
)

type Repository interface {
	CreateTrade(ctx context.Context, trade *TradeRequest) error
	GetTradesByUserID(ctx context.Context, userID uuid.UUID) ([]TradeRequest, error)
  UpdateTradeStatus(ctx context.Context, tradeID uuid.UUID, status string) error
  TradeExists(ctx context.Context, senderID, receiverID, senderSkillID, receiverSkillID uuid.UUID) (bool, error)
  checkSenderSkillOffering(senderID, senderSkillID uuid.UUID) (bool, error)
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

// TradeExists checks if a pending trade request already exists between two users for the same skill exchange.
func (r *repository) TradeExists(ctx context.Context, senderID, receiverID, senderSkillID, receiverSkillID uuid.UUID) (bool, error) {
    var count int64
    err := r.db.Model(&TradeRequest{}).
        Where("sender_id = ? AND receiver_id = ? AND sender_skill_id = ? AND receiver_skill_id = ? AND status = 'pending'",
            senderID, receiverID, senderSkillID, receiverSkillID).
        Count(&count).Error

    return count > 0, err
}

// checkSenderSkillOffering checks if the sender's skill is set as "offering" in the user_skills table
func (r *repository) checkSenderSkillOffering(senderID, senderSkillID uuid.UUID) (bool, error) {
    var userSkill struct {
        Type string `gorm:"column:type"`
    }
    err := r.db.Model(&userskills.UserSkill{}).
        Where("user_id = ? AND skill_id = ?", senderID, senderSkillID).
        Select("type").
        First(&userSkill).Error
    if err != nil {
        return false, err
    }

    // Check if the type is "offering"
    return userSkill.Type == "offering", nil
}
