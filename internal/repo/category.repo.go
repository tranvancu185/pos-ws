package repo

import (
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
)

type ICategoryRepo interface {
	GetInfo() string
}

type categoryRepo struct {
	sqlc *database.Queries
}

func NewCategoryRepo() ICategoryRepo {
	return &categoryRepo{
		sqlc: database.New(global.Mdbc),
	}
}

func (ur *categoryRepo) GetInfo() string {
	return "Hello"
}
