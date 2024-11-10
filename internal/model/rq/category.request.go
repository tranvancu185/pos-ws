package rq

type GetListCategoryRequest struct {
	PageSize  int64  `form:"page_size"`
	Page      int64  `form:"page"`
	Total     int64  `form:"total"`
	Name      string `form:"name"`
	Status    int64  `form:"status"`
	CreatedAt string `form:"created_at"`
	DeletedAt string `form:"deleted_at"`
}

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Status      int64  `json:"status"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int64  `json:"status"`
}
