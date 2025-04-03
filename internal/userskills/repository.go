package userskills

import (
	"errors"

	"gorm.io/gorm"
	"github.com/google/uuid"
  "github.com/GnotAI/skilltrade/internal/users"
  "github.com/GnotAI/skilltrade/internal/skills"
)

type UserSkillRepository struct {
	DB *gorm.DB
}

// NewUserSkillRepository creates a new repository instance
func NewUserSkillRepository(db *gorm.DB) *UserSkillRepository {
	return &UserSkillRepository{DB: db}
}

func (r *UserSkillRepository) CreateUserSkill(userID, skillID uuid.UUID, skillType string) error {
	// Ensure the skillType is either "offering" or "seeking"
	if skillType != "offering" && skillType != "seeking" {
		return errors.New("invalid skill type: must be 'offering' or 'seeking'")
	}

	// Check if user exists
	var user users.User
	if err := r.DB.First(&user, "id = ?", userID).Error; err != nil {
		return errors.New("user not found")
	}

	// Check if skill exists
	var skill skills.Skill
	if err := r.DB.First(&skill, "id = ?", skillID).Error; err != nil {
		return errors.New("skill not found")
	}

	// Create UserSkill entry
	userSkill := UserSkill{
		UserID:  userID,
		SkillID: skillID,
		Type:    skillType,
	}

	// Save the UserSkill entry to the database
	if err := r.DB.Create(&userSkill).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserSkillRepository) GetUserSkill(userID, skillID uuid.UUID, skillType string) (*UserSkill, error) {
    var userSkill *UserSkill
    err := r.DB.Preload("User").Preload("Skill").
        Where("user_id = ? AND skill_id = ? AND type = ?", userID, skillID, skillType).
        First(&userSkill).Error

    if err != nil {
        return nil, err
    }

    return userSkill, nil
}

// UserHasSkill checks if a user already has a specific skill
func (r *UserSkillRepository) UserHasSkill(userID, skillID uuid.UUID) (bool, error) {
	var count int64
	err := r.DB.Model(&UserSkill{}).
		Where("user_id = ? AND skill_id = ?", userID, skillID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
