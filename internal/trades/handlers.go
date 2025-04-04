package trades

import (
	"io"
	"net/http"

	jwtutil "github.com/GnotAI/skilltrade/utils/jwt"
	"github.com/bytedance/sonic"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler struct {
	Service *service
}

type UpdateTradeStatusPayload struct {
	Status string `json:"status"`
}

func NewTradeHandler(service *service) *Handler {
	return &Handler{Service: service}
}

type TradeRequestPayload struct {
	ReceiverID      uuid.UUID `json:"receiver_id"`
	SenderSkillID   uuid.UUID `json:"sender_skill_id"`
	ReceiverSkillID uuid.UUID `json:"receiver_skill_id"`
}

func (h *Handler) CreateTrade(w http.ResponseWriter, r *http.Request) {
  userID := r.Context().Value("user_id").(uuid.UUID)

  body, err := io.ReadAll(r.Body)
  if err != nil {
    http.Error(w, "invalid payload", http.StatusBadRequest)
    return
  }

  defer r.Body.Close()

	var payload TradeRequestPayload
	if err := sonic.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	trade, err := h.Service.RequestTrade(r.Context(), userID, payload.ReceiverID, payload.SenderSkillID, payload.ReceiverSkillID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	sonic.Marshal(trade)
}

func (h *Handler) GetMyTrades(w http.ResponseWriter, r *http.Request) {
  userID := r.Context().Value("user_id").(uuid.UUID)

	trades, err := h.Service.GetUserTrades(r.Context(), userID)
	if err != nil {
		http.Error(w, "Could not fetch trades", http.StatusInternalServerError)
		return
	}

	sonic.Marshal(trades)
}

func (h *Handler) UpdateTradeStatus(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uuid.UUID)

  body, err := io.ReadAll(r.Body)
  if err != nil {
    http.Error(w, "invalid payload", http.StatusBadRequest)
    return
  }

  defer r.Body.Close()

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

	err = h.Service.UpdateTradeStatus(r.Context(), tradeID, userID, payload.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
