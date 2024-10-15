package request

type GetListTableRequest struct {
	GetListRequest
	TableID     int64  `json:"table_id"`
	TableStatus int64  `json:"table_status"`
	Text        string `json:"text"`
}
