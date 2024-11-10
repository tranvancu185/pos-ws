package repo

import (
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/model/rq"
	"tranvancu185/vey-pos-ws/internal/uconst"
	"tranvancu185/vey-pos-ws/pkg/utils/utime"
)

type ICategoryRepo interface {
	GetListCategory(params *rq.GetListCategoryRequest) ([]database.GetListCategoriesRow, error)
	GetTotalCategory(params *rq.GetListCategoryRequest) (int64, error)
	GetCategoryByID(id int64) (*database.GetCategoryByIDRow, error)
	CreateCategory(params *rq.CreateCategoryRequest) (int64, error)
	UpdateCategoryByID(id int64, params *rq.UpdateCategoryRequest) error
	DeleteCategoryByID(id int64) error
}

type categoryRepo struct {
	sqlc *database.Queries
}

func NewCategoryRepo() ICategoryRepo {
	return &categoryRepo{
		sqlc: database.New(global.Mdbc),
	}
}

func (ur *categoryRepo) GetListCategory(params *rq.GetListCategoryRequest) ([]database.GetListCategoriesRow, error) {

	var input database.GetListCategoriesParams

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

	if params.Name != "" {
		input.CategoryName = params.Name
	}

	result, err := ur.sqlc.GetListCategories(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ur *categoryRepo) GetTotalCategory(params *rq.GetListCategoryRequest) (int64, error) {

	var input database.GetTotalCategoriesParams

	if params.Name != "" {
		input.CategoryName = params.Name
	}

	total, err := ur.sqlc.GetTotalCategories(ctx, input)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (ur *categoryRepo) GetCategoryByID(id int64) (*database.GetCategoryByIDRow, error) {
	result, err := ur.sqlc.GetCategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ur *categoryRepo) CreateCategory(params *rq.CreateCategoryRequest) (int64, error) {

	var input database.CreateCategoryParams
	currentTime := utime.GetCurrentTimeUnix()

	input.CategoryName = params.Name
	input.CreatedAt.Int64 = currentTime
	input.UpdatedAt.Int64 = currentTime

	id, err := ur.sqlc.CreateCategory(ctx, input)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ur *categoryRepo) UpdateCategoryByID(id int64, params *rq.UpdateCategoryRequest) error {
	var input database.UpdateCategoryByIDParams
	currentTime := utime.GetCurrentTimeUnix()

	input.CategoryID = id
	input.CategoryName = params.Name
	input.UpdatedAt.Int64 = currentTime

	err := ur.sqlc.UpdateCategoryByID(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (ur *categoryRepo) DeleteCategoryByID(id int64) error {
	var input database.DeleteCategoryByIDParams
	currentTime := utime.GetCurrentTimeUnix()

	input.CategoryID = id
	input.DeletedAt.Int64 = currentTime

	err := ur.sqlc.DeleteCategoryByID(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
