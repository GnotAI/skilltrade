package sessions

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/GnotAI/skilltrade/internal/trades"
)

type service struct {
	Repo          *sessRepository
	tradeRepo     *trades.TradeRepository
}

func NewService(repo *sessRepository, tradeRepo *trades.TradeRepository) *service {
	return &service{Repo: repo, tradeRepo: tradeRepo}
}

func (s *service) ScheduleSession(userID, tradeID uuid.UUID, scheduledAt time.Time) (*Session, error) {

	// Check if the trade exists
	trade, err := s.tradeRepo.GetTradeRequestByID(tradeID)
	if err != nil {
		return nil, errors.New("trade request not found")
	}

	// Ensure the user is either the sender or receiver of the trade
	if trade.SenderID != userID && trade.ReceiverID != userID {
		return nil, errors.New("user is not authorized to schedule a session for this trade")
	}

	// Ensure the trade request is accepted before scheduling a session
	if trade.Status != "accepted" {
		return nil, errors.New("cannot schedule a session for a trade that is not accepted")
	}

	// Check if the scheduled time is in the future
	if scheduledAt.Before(time.Now()) {
		return nil, errors.New("scheduled time must be in the future")
	}

	// Create the session
	session := &Session{
		TradeID:     trade.ID,
		ScheduledAt: scheduledAt,
	}

	if err := s.Repo.CreateSession(session); err != nil {
		return nil, err
	}

  return session, nil
}

func (s *service) MarkSessionCompleted(userID, sessionID uuid.UUID) error {
	// Retrieve the session with the trade relationship
	var session Session
	err := s.Repo.DB.Preload("Trade").First(&session, "id = ?", sessionID).Error
	if err != nil {
		return errors.New("session not found")
	}

	// Check if the user is the sender or receiver in the related trade
	if session.Trade.SenderID != userID && session.Trade.ReceiverID != userID {
		return errors.New("user is not authorized to mark this session as completed")
	}

	// Update the session as completed
	return s.Repo.MarkSessionCompleted(sessionID)
}
