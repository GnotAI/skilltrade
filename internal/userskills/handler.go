package userskills

import (
	"io"
	"net/http"

	"github.com/GnotAI/skilltrade/utils/jwt"
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
  token, ok := r.Context().Value("AuthorizationToken").(string)
  if !ok {
		http.Error(w, "Authorization token not found", http.StatusUnauthorized)
		return
  }

	claims, err := jwtutil.ParseToken(token)
	if err != nil {
		http.Error(w, "invalid token", http.StatusBadRequest)
    return
	}

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

  claimsID, err := uuid.Parse(claims.UserID)

  if claimsID != req.UserID {
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
