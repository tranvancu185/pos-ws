package service

import (
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/model/rq"
	"tranvancu185/vey-pos-ws/internal/repo"
)

type IProductService interface {
	GetListProduct(params *rq.GetListProductRequest) ([]database.GetListProductsRow, error)
	SearchProduct(params *rq.SearchProductRequest) ([]database.SearchProductsRow, error)
	CreateProduct(params *rq.CreateProductRequest) (int64, error)
	UpdateProduct(id int64, params *rq.UpdateProductRequest) error
	GetProductByID(id int64) (*database.GetProductByIDRow, error)
	GetTotalProduct(params *rq.GetListProductRequest) (int64, error)
	DeleteProductByID(id int64) error
	UpdateProductImageByID(id int64, image string) error
}

type productService struct {
	productRepo repo.IProductRepo
}

func NewProductService(productRepo repo.IProductRepo) IProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (ps *productService) GetListProduct(params *rq.GetListProductRequest) ([]database.GetListProductsRow, error) {
	products, err := ps.productRepo.GetListProduct(params)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *productService) SearchProduct(params *rq.SearchProductRequest) ([]database.SearchProductsRow, error) {
	products, err := ps.productRepo.SearchProduct(params)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *productService) CreateProduct(params *rq.CreateProductRequest) (int64, error) {
	code, errCode := NewCommonService().GenerateCode(COUNTER_PRODUCT)
	if errCode != nil {
		return 0, errCode
	}

	params.ProductCode = code

	id, err := ps.productRepo.CreateProduct(params)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ps *productService) UpdateProduct(id int64, params *rq.UpdateProductRequest) error {
	err := ps.productRepo.UpdateProductByID(id, params)
	if err != nil {
		return err
	}
	return nil
}

func (ps *productService) GetProductByID(id int64) (*database.GetProductByIDRow, error) {
	product, err := ps.productRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *productService) GetTotalProduct(params *rq.GetListProductRequest) (int64, error) {
	total, err := ps.productRepo.GetTotalProduct(params)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (ps *productService) DeleteProductByID(id int64) error {
	err := ps.productRepo.DeleteProductByID(id)
	if err != nil {
		return err
	}
	return nil
}

func (ps *productService) UpdateProductImageByID(id int64, image string) error {
	err := ps.productRepo.UpdateProductImageByID(id, image)
	if err != nil {
		return err
	}
	return nil
}
