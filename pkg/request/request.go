package request

type GetListRequest struct {
	PageSize int64 `json:"page_size"`
	Page     int64 `json:"page"`
	Total    int64 `json:"total"`
}
