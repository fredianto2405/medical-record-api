package menu

type DTO struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Icon       string `json:"icon" db:"icon"`
	SortNumber int    `json:"sort_number" db:"sort_number"`
	PathName   string `json:"path_name" db:"path_name"`
	ParentID   int    `json:"parent_id" db:"parent_id"`
	Children   []*DTO `json:"children"`
}
