package trades

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	RequestTrade(ctx context.Context, senderID, receiverID, senderSkillID, receiverSkillID uuid.UUID) (*TradeRequest, error)
	GetUserTrades(ctx context.Context, userID uuid.UUID) ([]TradeRequest, error)
  UpdateTradeStatus(ctx context.Context, tradeID uuid.UUID, userID uuid.UUID, status string) error
}

type service struct {
	repo Repository
}

func NewTradeService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) RequestTrade(ctx context.Context, senderID, receiverID, senderSkillID, receiverSkillID uuid.UUID) (*TradeRequest, error) {
	if senderID == receiverID {
		return nil, errors.New("cannot trade with yourself")
	}

	trade := &TradeRequest{
		SenderID:        senderID,
		ReceiverID:      receiverID,
		SenderSkillID:   senderSkillID,
		ReceiverSkillID: receiverSkillID,
	}

	err := s.repo.CreateTrade(ctx, trade)
	return trade, err
}

func (s *service) GetUserTrades(ctx context.Context, userID uuid.UUID) ([]TradeRequest, error) {
	return s.repo.GetTradesByUserID(ctx, userID)
}

func (s *service) UpdateTradeStatus(ctx context.Context, tradeID uuid.UUID, userID uuid.UUID, status string) error {
	if status != "accepted" && status != "rejected" {
		return errors.New("invalid status: must be 'accepted' or 'rejected'")
	}

	// Optional: validate that the trade exists and belongs to the receiver
	trades, err := s.repo.GetTradesByUserID(ctx, userID)
	if err != nil {
		return err
	}

	found := false
	for _, t := range trades {
		if t.ID == tradeID && t.ReceiverID == userID {
			found = true
			break
		}
	}
	if !found {
		return errors.New("unauthorized or trade not found")
	}

	return s.repo.UpdateTradeStatus(ctx, tradeID, status)
}
