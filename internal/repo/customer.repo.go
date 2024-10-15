package repo

import (
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
)

type ICustomerRepo interface {
	GetInfo() string
}

type customerRepo struct {
	sqlc *database.Queries
}

func NewCustomerRepo() ICustomerRepo {
	return &customerRepo{
		sqlc: database.New(global.Mdbc),
	}
}

func (ur *customerRepo) GetInfo() string {
	return "Hello"
}
