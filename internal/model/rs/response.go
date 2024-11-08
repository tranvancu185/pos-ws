package rs

type GetListResponse struct {
	Page     int64       `json:"page"`
	PageSize int64       `json:"page_size"`
	Total    int64       `json:"total"`
	Rows     interface{} `json:"rows"`
}
