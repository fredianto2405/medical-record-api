package medicine

import (
	mapper "medical-record-api/internal/medicine/mapper"
	model "medical-record-api/internal/medicine/model"
	repository "medical-record-api/internal/medicine/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo}
}

func (s *CategoryService) GetAll(search string) ([]*model.CategoryDTO, error) {
	return s.repo.FindAll(search)
}

func (s *CategoryService) GetAllPaginated(page, limit int, search string) ([]*model.CategoryDTO, int, error) {
	return s.repo.FindAllPaginated(page, limit, search)
}

func (s *CategoryService) Create(request *model.CategoryRequest) (*model.CategoryDTO, error) {
	entity := mapper.MapToCategoryEntity(request)

	savedEntity, err := s.repo.Save(entity)
	if err != nil {
		return nil, err
	}

	return mapper.MapToCategoryDTO(savedEntity), nil
}

func (s *CategoryService) Update(id string, request *model.CategoryRequest) (*model.CategoryDTO, error) {
	entity := mapper.MapToCategoryEntity(request)
	entity.ID = id

	updatedEntity, err := s.repo.Update(entity)
	if err != nil {
		return nil, err
	}

	return mapper.MapToCategoryDTO(updatedEntity), nil
}

func (s *CategoryService) Delete(id string) error {
	return s.repo.Delete(id)
}
