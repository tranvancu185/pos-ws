package repo

import (
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/model/rq"
	"tranvancu185/vey-pos-ws/internal/uconst"
	"tranvancu185/vey-pos-ws/pkg/utils/utime"
)

type IProductRepo interface {
	GetListProduct(params *rq.GetListProductRequest) ([]database.GetListProductsRow, error)
	GetTotalProduct(params *rq.GetListProductRequest) (int64, error)
	GetProductByID(id int64) (*database.GetProductByIDRow, error)
	CreateProduct(params *rq.CreateProductRequest) (int64, error)
	UpdateProductByID(id int64, params *rq.UpdateProductRequest) error
	DeleteProductByID(id int64) error
	SearchProduct(params *rq.SearchProductRequest) ([]database.SearchProductsRow, error)
	GetProductByCode(code string) (*database.GetProductByCodeRow, error)
	UpdateProductImageByID(id int64, image string) error
}

type productRepo struct {
	sqlc *database.Queries
}

func NewProductRepo() IProductRepo {
	return &productRepo{
		sqlc: database.New(global.Mdbc),
	}
}

func (pr *productRepo) GetListProduct(params *rq.GetListProductRequest) ([]database.GetListProductsRow, error) {
	var input database.GetListProductsParams

	if params.PageSize != 0 {
		input.Limit = params.PageSize
	} else {
		input.Limit = uconst.DEFAULT_LIMIT
	}

	if params.Page != 0 {
		input.Offset = (params.Page - 1) * params.PageSize
	} else {
		input.Offset = uconst.DEFAULT_OFFSET
	}

	if params.Text != "" {
		input.ProductCode = "%" + params.Text + "%"
		input.ProductName = "%" + params.Text + "%"
	}

	if params.ProductStatus != 0 {
		input.ProductStatus = params.ProductStatus
	}

	if params.ProductCategoryId != 0 {
		input.ProductCategoryID = params.ProductCategoryId
	}

	result, err := pr.sqlc.GetListProducts(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pr *productRepo) GetTotalProduct(params *rq.GetListProductRequest) (int64, error) {
	var input database.GetTotalProductsParams

	if params.ProductStatus != 0 {
		input.ProductStatus = params.ProductStatus
	}

	total, err := pr.sqlc.GetTotalProducts(ctx, input)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (pr *productRepo) GetProductByID(id int64) (*database.GetProductByIDRow, error) {
	product, err := pr.sqlc.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (pr *productRepo) CreateProduct(params *rq.CreateProductRequest) (int64, error) {
	var input database.CreateProductParams
	currentTime := utime.GetCurrentTimeUnix()

	input.ProductCode = params.ProductCode
	input.ProductName = params.ProductName
	input.ProductPrice = params.ProductPrice
	input.ProductDisplayName = params.ProductDisplayName
	input.ProductDescription.String = params.ProductDescription
	input.ProductStatus = params.ProductStatus
	input.ProductCategoryID = params.ProductCategoryId
	input.CreatedAt.Int64 = currentTime
	input.UpdatedAt.Int64 = currentTime

	id, err := pr.sqlc.CreateProduct(ctx, input)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (pr *productRepo) UpdateProductByID(id int64, params *rq.UpdateProductRequest) error {
	var input database.UpdateProductByIDParams

	if params.ProductName != "" {
		input.ProductName = params.ProductName
	}

	if params.ProductStatus != 0 {
		input.ProductStatus = params.ProductStatus
	}

	if params.ProductPrice != 0 {
		input.ProductPrice = params.ProductPrice
	}

	if params.ProductDisplayName != "" {
		input.ProductDisplayName = params.ProductDisplayName
	}

	if params.ProductDescription != "" {
		input.ProductDescription.String = params.ProductDescription
	}

	if params.ProductCategoryId != 0 {
		input.ProductCategoryID = params.ProductCategoryId
	}

	err := pr.sqlc.UpdateProductByID(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (pr *productRepo) DeleteProductByID(id int64) error {
	var input database.DeleteProductByIDParams
	currentTime := utime.GetCurrentTimeUnix()

	input.ProductID = id
	input.DeletedAt.Int64 = currentTime

	err := pr.sqlc.DeleteProductByID(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (pr *productRepo) SearchProduct(params *rq.SearchProductRequest) ([]database.SearchProductsRow, error) {
	var input database.SearchProductsParams

	if params.Text != "" {
		input.ProductName = "%" + params.Text + "%"
		input.ProductCode = "%" + params.Text + "%"
	}

	input.ProductStatus = 1

	result, err := pr.sqlc.SearchProducts(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pr *productRepo) GetProductByCode(code string) (*database.GetProductByCodeRow, error) {
	product, err := pr.sqlc.GetProductByCode(ctx, code)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (pr *productRepo) UpdateProductImageByID(id int64, image string) error {
	var input database.UpdateProductImageByIDParams

	input.ProductID = id
	input.ProductImage = image
	input.UpdatedAt.Int64 = utime.GetCurrentTimeUnix()

	err := pr.sqlc.UpdateProductImageByID(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
