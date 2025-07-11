package treatment

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) GetAll(search string) ([]*DTO, error) {
	return s.repo.FindAll(search)
}

func (s *Service) GetAllPaginated(page, limit int, search string) ([]*DTO, int, error) {
	return s.repo.FindAllPaginated(page, limit, search)
}

func (s *Service) Create(request *Request) (*DTO, error) {
	entity := MapToEntity(request)

	savedEntity, err := s.repo.Save(entity)
	if err != nil {
		return nil, err
	}

	return MapToDTO(savedEntity), nil
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

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
