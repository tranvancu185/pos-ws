package request

type GetListTableRequest struct {
	GetListRequest
	TableID     int64  `json:"table_id"`
	TableStatus int64  `json:"table_status"`
	Text        string `json:"text"`
}

type CreateTableRequest struct {
	TableName string `json:"table_name" validate:"required"`
}

type UpdateTableRequest struct {
	TableID     int64  `json:"table_id" validate:"required"`
	TableName   string `json:"table_name"`
	TableStatus int64  `json:"table_status"`
}
