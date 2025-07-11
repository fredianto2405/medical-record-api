package menu

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindByRoleID(roleID int) ([]*DTO, error) {
	var menus []*DTO

	dataQuery := `select m.id,
			m.name,
			m.icon, 
			m.sort_number,
			m.path_name,
			m.parent_id
		from emr_menu.menus m
		join emr_menu.menu_permissions mp
		on mp.menu_id = m.id
		where m.is_active = true
		and mp.is_active = true
		and mp.role_id = $1
		order by m.sort_number asc`

	err := r.db.Select(&menus, dataQuery, roleID)
	return menus, err
}
