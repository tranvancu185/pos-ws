package service

import "tranvancu185/vey-pos-ws/internal/repo"

type IProductService interface {
	GetInFo() string
}

type productService struct {
	productRepo repo.IProductRepo
}

func NewProductService(productRepo repo.IProductRepo) IProductService {
	return &productService{
		productRepo: productRepo,
	}
}

// GetInFo implements IUserService.
func (prs *productService) GetInFo() string {
	return (prs.productRepo).GetInFo()
}
