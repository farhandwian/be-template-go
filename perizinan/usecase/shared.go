package usecase

type Metadata struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}
