package rq

type GetListRequest struct {
	PageSize int64 `form:"page_size"`
	Page     int64 `form:"page"`
	Total    int64 `form:"total"`
}
