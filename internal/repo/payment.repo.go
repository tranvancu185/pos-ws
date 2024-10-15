package repo

import (
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
)

type IPaymentRepo interface {
	GetInfo() string
}

type paymentRepo struct {
	sqlc *database.Queries
}

func NewPaymentRepo() IPaymentRepo {
	return &paymentRepo{
		sqlc: database.New(global.Mdbc),
	}
}

func (ur *paymentRepo) GetInfo() string {
	return "Hello"
}
