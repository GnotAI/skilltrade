package skills

import (
  "errors"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type SkillRepository struct {
	DB *gorm.DB
}

func NewSkillRepository(db *gorm.DB) *SkillRepository {
	return &SkillRepository{DB: db}
}


// GetAllSkills retrieves all skills from the database
func (r *SkillRepository) GetAllSkills() ([]Skill, error) {
	var skills []Skill
	if err := r.DB.Find(&skills).Error; err != nil {
		return nil, err
	}
	return skills, nil
}

func (r *SkillRepository) GetSkillByDescription(description string) (*Skill, error) {
	var skill Skill
	if err := r.DB.Where("description = ?", description).First(&skill).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("skill not found")
		}
		return nil, err
	}
	return &skill, nil
}

func (r *SkillRepository) GetSkillByID(ID uuid.UUID) (*Skill, error) {
    var skill Skill
    result := r.DB.First(&skill, "id = ?", ID)
    if result.Error != nil {
        return nil, result.Error
    }
    return &skill, nil
}

func (r *SkillRepository) CreateSkill(skill *Skill) error {
	return r.DB.Create(skill).Error
}

func (r *SkillRepository) UpdateSkill(skill *Skill) error {
	return r.DB.Save(skill).Error
}

func (r *SkillRepository) DeleteSkill(id uuid.UUID) error {
	return r.DB.Delete(&Skill{}, "id = ?", id).Error
}
