package trades

import (
	"context"
	"errors"
	"fmt"

	"github.com/GnotAI/skilltrade/internal/userskills"
	"github.com/google/uuid"
)

type service struct {
	repo *TradeRepository
	UserSkillService *userskills.UserSkillService
}

func NewTradeService(repo *TradeRepository) *service {
	return &service{repo: repo}
}

func (s *service) RequestTrade(ctx context.Context, senderID, receiverID, senderSkillID, receiverSkillID uuid.UUID) (*TradeRequest, error) {
  if senderID == receiverID {
    return nil, errors.New("cannot trade with yourself")
  }

  if senderSkillID == receiverSkillID {
    return nil, errors.New("cannot trade with the same skill")
  }

  // Check if the sender and receiver skills are marked as "offering" in their user_skills entries
  isSendOffering, err := s.repo.checkSenderSkillOffering(senderID, senderSkillID)
  if err != nil {
    return nil, fmt.Errorf("failed to check sender's skill status: %w", err)
  }
  if !isSendOffering {
    return nil, errors.New("sender's skill is not marked as 'offering'")
  }

  isRecOffering, err := s.repo.checkSenderSkillOffering(receiverID, receiverSkillID)
  if err != nil {
    return nil, fmt.Errorf("failed to check receiver's skill status: %w", err)
  }
  if !isRecOffering {
    return nil, errors.New("receiver's skill is not marked as 'offering'")
  }

  // Check if a similar pending trade request exists
  exists, err := s.repo.TradeExists(ctx, senderID, receiverID, senderSkillID, receiverSkillID)
  if err != nil {
    return nil, fmt.Errorf("failed to check existing trades: %w", err)
  }
  if exists {
    return nil, errors.New("a pending trade request already exists between these users for the same skills")
  }


  trade := &TradeRequest{
    SenderID:        senderID,
    ReceiverID:      receiverID,
    SenderSkillID:   senderSkillID,
    ReceiverSkillID: receiverSkillID,
  }

  err = s.repo.CreateTrade(ctx, trade)
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
