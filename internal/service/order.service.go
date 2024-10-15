package service

import "tranvancu185/vey-pos-ws/internal/repo"

type IOrderService interface {
	GetInfo() string
}

type orderService struct {
	orderRepo repo.IOrderRepo
}

func NewOrderService(orderRepo repo.IOrderRepo) IOrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

// GetInfo implements IOrder.
func (os *orderService) GetInfo() string {
	return (os.orderRepo).GetInfo()
}
