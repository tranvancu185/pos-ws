package service

import (
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/model/rq"
	"tranvancu185/vey-pos-ws/internal/repo"
)

type ICategoryService interface {
	GetListCategories(params *rq.GetListCategoryRequest) ([]database.GetListCategoriesRow, error)
	CreateCategory(params *rq.CreateCategoryRequest) (int64, error)

	UpdateCategory(id int64, params *rq.UpdateCategoryRequest) error
	GetCategoryByID(id int64) (*database.GetCategoryByIDRow, error)
	GetTotalCategories(params *rq.GetListCategoryRequest) (int64, error)
	DeleteCategoryByID(id int64) error
}

type categoryService struct {
	categoryRepo repo.ICategoryRepo
}

func NewCategoryService(categoryRepo repo.ICategoryRepo) ICategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (cs *categoryService) GetListCategories(params *rq.GetListCategoryRequest) ([]database.GetListCategoriesRow, error) {
	categories, err := cs.categoryRepo.GetListCategory(params)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (cs *categoryService) CreateCategory(params *rq.CreateCategoryRequest) (int64, error) {
	id, err := cs.categoryRepo.CreateCategory(params)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (cs *categoryService) UpdateCategory(id int64, params *rq.UpdateCategoryRequest) error {
	err := cs.categoryRepo.UpdateCategoryByID(id, params)
	if err != nil {
		return err
	}
	return nil
}

func (cs *categoryService) GetCategoryByID(id int64) (*database.GetCategoryByIDRow, error) {
	category, err := cs.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (cs *categoryService) GetTotalCategories(params *rq.GetListCategoryRequest) (int64, error) {
	total, err := cs.categoryRepo.GetTotalCategory(params)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (cs *categoryService) DeleteCategoryByID(id int64) error {
	err := cs.categoryRepo.DeleteCategoryByID(id)
	if err != nil {
		return err
	}
	return nil
}
