package rq

type GetListProductRequest struct {
	PageSize          int64  `form:"page_size"`
	Page              int64  `form:"page"`
	Total             int64  `form:"total"`
	Text              string `form:"text"`
	CreateAt          string `form:"create_at"`
	ProductCategoryId int64  `form:"product_category_id"`
	ProductStatus     int64  `form:"product_status"`
}

type SearchProductRequest struct {
	Text string `form:"text"`
}

type CreateProductRequest struct {
	ProductCode        string `json:"product_code"`
	ProductName        string `json:"product_name" validate:"required"`
	ProductStatus      int64  `json:"product_status"`
	ProductCategoryId  int64  `json:"product_category_id" validate:"required"`
	ProductDisplayName string `json:"product_display_name" validate:"required"`
	ProductDescription string `json:"product_description" validate:"required"`
	ProductPrice       int64  `json:"product_price" validate:"required"`
	ProductImage       string `json:"product_image"`
}

type UpdateProductRequest struct {
	ProductCode        string `json:"product_code"`
	ProductName        string `json:"product_name"`
	ProductStatus      int64  `json:"product_status"`
	ProductPrice       int64  `json:"product_price"`
	ProductDisplayName string `json:"product_display_name"`
	ProductDescription string `json:"product_description"`
	ProductCategoryId  int64  `json:"product_category_id"`
}
