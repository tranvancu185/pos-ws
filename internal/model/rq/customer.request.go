package rq

type GetListCustomerRequest struct {
	PageSize       int64  `form:"page_size"`
	Page           int64  `form:"page"`
	Total          int64  `form:"total"`
	Text           string `form:"text"`
	CustomerStatus int64  `form:"customer_status"`
}

type CreateCustomerRequest struct {
	CustomerName  string `json:"customer_name" validate:"required"`
	CustomerPhone string `json:"customer_phone" validate:"required"`
	CustomerCode  string `json:"customer_code"`
	CustomerEmail string `json:"customer_email"`
}

type UpdateCustomerRequest struct {
	CustomerName   string `json:"customer_name"`
	CustomerPhone  string `json:"customer_phone"`
	CustomerEmail  string `json:"customer_email"`
	CustomerStatus int64  `json:"customer_status"`
	Properties     string `json:"properties"`
}
