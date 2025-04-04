package sessions

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bytedance/sonic"
	"github.com/google/uuid"
	"github.com/go-chi/chi/v5"
)

type SessionHandler struct {
	Service *service
}

func NewSessionHandler(service *service) *SessionHandler {
	return &SessionHandler{Service: service}
}

// ScheduleSession handler for scheduling a session.
func (h *SessionHandler) ScheduleSession(w http.ResponseWriter, r *http.Request) {
  userID, ok := r.Context().Value("user_id").(uuid.UUID)
  if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
  }

	// Parse request body to get trade ID and scheduled time
	var req struct {
		TradeID     uuid.UUID `json:"trade_id"`
		ScheduledAt time.Time `json:"scheduled_at"`
	}

  body, err := io.ReadAll(r.Body)
  if err != nil {
    http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
    return
  }
  defer r.Body.Close()

	if err := sonic.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the service to schedule the session
	session, err := h.Service.ScheduleSession(userID, req.TradeID, req.ScheduledAt)
	if err != nil {
		// Handle specific errors
		if err.Error() == "trade request not found" {
			http.Error(w, "Trade request not found", http.StatusNotFound)
		} else if err.Error() == "cannot schedule a session for a trade that is not accepted" {
			http.Error(w, "Cannot schedule a session for a trade that is not accepted", http.StatusConflict)
		} else if err.Error() == "scheduled time must be in the future" {
			http.Error(w, "Scheduled time must be in the future", http.StatusBadRequest)
		} else {
			http.Error(w, "Failed to schedule session", http.StatusInternalServerError)
		}
		return
	}

	// Return the created session as a response
	w.WriteHeader(http.StatusCreated)
  _, err = sonic.Marshal(session)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *SessionHandler) MarkSessionCompleted(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)

  ID, err := uuid.Parse(userID)
  if err != nil {
    http.Error(w, fmt.Sprintf("%s", err), http.StatusUnauthorized)
  }

	sessionIDStr := chi.URLParam(r, "id")
	sessionID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		http.Error(w, "invalid session ID", http.StatusBadRequest)
		return
	}

	err = h.Service.MarkSessionCompleted(ID, sessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("session marked as completed"))
}
