package menu

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetByRoleID(roleID int) ([]*DTO, error) {
	menuMap := make(map[int]*DTO)
	var roots []*DTO

	menus, err := s.repo.FindByRoleID(roleID)
	if err != nil {
		return nil, err
	}

	// convert to map
	for i := range menus {
		menu := menus[i]
		menu.Children = []*DTO{}
		menuMap[menu.ID] = menu
	}

	// build tree
	for _, menu := range menus {
		if menu.ParentID == 0 {
			roots = append(roots, menu)
		} else {
			if parent, exists := menuMap[menu.ParentID]; exists {
				parent.Children = append(parent.Children, menu)
			}
		}
	}

	return roots, nil
}
