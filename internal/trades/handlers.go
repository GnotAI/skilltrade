package trades

import (
	"io"
	"net/http"

	jwtutil "github.com/GnotAI/skilltrade/utils/jwt"
	"github.com/go-chi/chi/v5"
	"github.com/bytedance/sonic"
	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

type UpdateTradeStatusPayload struct {
	Status string `json:"status"`
}

func NewTradeHandler(service Service) *Handler {
	return &Handler{service: service}
}

type TradeRequestPayload struct {
	ReceiverID      uuid.UUID `json:"receiver_id"`
	SenderSkillID   uuid.UUID `json:"sender_skill_id"`
	ReceiverSkillID uuid.UUID `json:"receiver_skill_id"`
}

func (h *Handler) CreateTrade(w http.ResponseWriter, r *http.Request) {
	claims, err := jwtutil.ParseToken(r.Context().Value("AuthorizationToken").(string))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

  body, err := io.ReadAll(r.Body)
  if err != nil {
    http.Error(w, "invalid payload", http.StatusBadRequest)
    return
  }

	var payload TradeRequestPayload
	if err := sonic.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

  claimsID, err := uuid.Parse(claims.UserID)
  if err != nil {
    http.Error(w, "Unauthorized", http.StatusUnauthorized)
    return
  }

	trade, err := h.service.RequestTrade(r.Context(), claimsID, payload.ReceiverID, payload.SenderSkillID, payload.ReceiverSkillID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	sonic.Marshal(trade)
}

func (h *Handler) GetMyTrades(w http.ResponseWriter, r *http.Request) {
	claims, err := jwtutil.ParseToken(r.Context().Value("AuthorizationToken").(string))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

  claimsID, err := uuid.Parse(claims.UserID)
  if err != nil {
    http.Error(w, "Unauthorized", http.StatusUnauthorized)
    return
  }

	trades, err := h.service.GetUserTrades(r.Context(), claimsID)
	if err != nil {
		http.Error(w, "Could not fetch trades", http.StatusInternalServerError)
		return
	}

	sonic.Marshal(trades)
}

func (h *Handler) UpdateTradeStatus(w http.ResponseWriter, r *http.Request) {
	claims, err := jwtutil.ParseToken(r.Context().Value("AuthorizationToken").(string))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

  body, err := io.ReadAll(r.Body)
  if err != nil {
    http.Error(w, "invalid payload", http.StatusBadRequest)
    return
  }

	tradeIDStr := chi.URLParam(r, "id")
	tradeID, err := uuid.Parse(tradeIDStr)
	if err != nil {
		http.Error(w, "Invalid trade ID", http.StatusBadRequest)
		return
	}

	var payload UpdateTradeStatusPayload
	if err := sonic.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

  claimsID, err := uuid.Parse(claims.UserID)
  if err != nil {
    http.Error(w, "Unauthorized", http.StatusUnauthorized)
    return
  }

	err = h.service.UpdateTradeStatus(r.Context(), tradeID, claimsID, payload.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
