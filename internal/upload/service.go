package upload

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(e *Entity) (string, error) {
	return s.repo.Save(e)
}
