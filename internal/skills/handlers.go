package skills

import (
	"github.com/bytedance/sonic"
	"net/http"
)

type SkillHandler struct {
	Service *SkillService
}

func NewSkillHandler(service *SkillService) *SkillHandler {
	return &SkillHandler{Service: service}
}

// GetAllSkillsHandler handles the request to retrieve all skills
func (h *SkillHandler) GetAllSkillsHandler(w http.ResponseWriter, r *http.Request) {
	skills, err := h.Service.GetAllSkills()
	if err != nil {
		http.Error(w, "Failed to retrieve skills", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := sonic.Marshal(skills)
	w.Write(jsonData)
}
