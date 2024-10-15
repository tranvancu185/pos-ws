package service

import "tranvancu185/vey-pos-ws/internal/repo"

type ICategoryService interface {
	GetInfo() string
}

type categoryService struct {
	categoryRepo repo.ICategoryRepo
}

func NewCategoryService(categoryRepo repo.ICategoryRepo) ICategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

// GetInfo implements ICategory.
func (cas *categoryService) GetInfo() string {
	return (cas.categoryRepo).GetInfo()
}
