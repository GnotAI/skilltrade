package userskills

import (
	"errors"

	"github.com/google/uuid"
)

type UserSkillService struct {
	Repo *UserSkillRepository
}

func NewUserSkillService(repo *UserSkillRepository) *UserSkillService {
	return &UserSkillService{Repo: repo}
}

func (s *UserSkillService) CreateUserSkillService(userID, skillID uuid.UUID, skillType string) error {
  
  // Ensure there is a user id
  if userID.String() == "" {
		return errors.New("UserID needed")
  }

  // Ensure there is a skill id
  if skillID.String() == "" {
		return errors.New("UserID needed")
  }

	// Ensure skillType is valid
	if skillType != "offering" && skillType != "seeking" {
		return errors.New("invalid skill type: must be 'offering' or 'seeking'")
	}

	// Check if the user already has this skill
	exists, err := s.Repo.UserHasSkill(userID, skillID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user already has this skill")
	}

	// Call the repository function to create the user skill
	err = s.Repo.CreateUserSkill(userID, skillID, skillType)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserSkillService) GetUserSkill(userID, skillID uuid.UUID) (*UserSkill, error) {
    return s.Repo.GetUserSkill(userID, skillID)
}
