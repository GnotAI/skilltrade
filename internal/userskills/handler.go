package userskills

import (
	"io"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/google/uuid"
)

type UserSkillHandler struct {
	Service *UserSkillService
}

func NewUserSkillHandler(service *UserSkillService) *UserSkillHandler {
	return &UserSkillHandler{Service: service}
}

func (h *UserSkillHandler) CreateUserSkillHandler(w http.ResponseWriter, r *http.Request) {
  userID := r.Context().Value("user_id").(uuid.UUID)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

  defer r.Body.Close()

  var req struct {
    UserID  uuid.UUID `json:"user_id"`
    SkillID uuid.UUID `json:"skill_id"`
    Type    string    `json:"type"`
  }

  if err := sonic.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
  }

  if userID != req.UserID {
		http.Error(w, "Please sign in", http.StatusUnauthorized)
		return
  }

	err = h.Service.CreateUserSkillService(req.UserID, req.SkillID, req.Type)
	if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, _ := sonic.Marshal(map[string]string{"message": "User skill added successfully"})
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
