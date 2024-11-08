package rq

type GetListTableRequest struct {
	PageSize    int64  `form:"page_size"`
	Page        int64  `form:"page"`
	Total       int64  `form:"total"`
	TableID     int64  `form:"table_id"`
	TableStatus int64  `form:"table_status"`
	Text        string `form:"text"`
}

type CreateTableRequest struct {
	TableName string `json:"table_name" validate:"required"`
	TableCode string `json:"table_code"`
}

type UpdateTableRequest struct {
	TableName   string `json:"table_name"`
	TableStatus int64  `json:"table_status"`
}
