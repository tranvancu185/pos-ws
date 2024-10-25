package request

type GetListTableRequest struct {
	GetListRequest
	TableID     int64  `json:"table_id"`
	TableStatus int64  `json:"table_status"`
	Text        string `json:"text"`
}

type CreateTableRequest struct {
	TableCode string `json:"table_code"`
	TableName string `json:"table_name"`
}

type UpdateTableRequest struct {
	TableID     int64  `json:"table_id"`
	TableCode   string `json:"table_code"`
	TableName   string `json:"table_name"`
	TableStatus int64  `json:"table_status"`
}
