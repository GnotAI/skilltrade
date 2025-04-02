package skills

type SkillService struct {
	Repo *SkillRepository
}

func NewSkillService(repo *SkillRepository) *SkillService {
	return &SkillService{Repo: repo}
}

// GetAllSkills retrieves all skills from the repository
func (s *SkillService) GetAllSkills() ([]Skill, error) {
	return s.Repo.GetAllSkills()
}
