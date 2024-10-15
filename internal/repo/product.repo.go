package repo

import (
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
)

type IProductRepo interface {
	GetInFo() string
}

type productRepo struct {
	sqlc *database.Queries
}

func NewProductRepo() IProductRepo {
	return &productRepo{
		sqlc: database.New(global.Mdbc),
	}
}

func (ur *productRepo) GetInFo() string {
	return "Hello"
}
