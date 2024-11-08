package repo

import (
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
)

type IOrderRepo interface {
	GetInfo() string
}

type orderRepo struct {
	sqlc *database.Queries
}

func NewOrderRepo() IOrderRepo {
	return &orderRepo{
		sqlc: database.New(global.Mdbc),
	}
}

func (ur *orderRepo) GetInfo() string {
	return "Hello"
}
