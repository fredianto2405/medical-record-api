package response

type Pagination struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
