package clinic

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) Get() (*DTO, error) {
	return s.repo.FindOne()
}

func (s *Service) Update(id string, request *Request) (*DTO, error) {
	entity := MapToEntity(request)
	entity.ID = id

	updatedEntity, err := s.repo.Update(entity)
	if err != nil {
		return nil, err
	}

	return MapToDTO(updatedEntity), nil
}
