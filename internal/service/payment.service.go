package service

import "tranvancu185/vey-pos-ws/internal/repo"

type IPaymentService interface {
	GetInfo() string
}

type paymentService struct {
	paymentRepo repo.IPaymentRepo
}

func NewPaymentService(paymentRepo repo.IPaymentRepo) IPaymentService {
	return &paymentService{
		paymentRepo: paymentRepo,
	}
}

// GetInfo implements IPayment.
func (ps *paymentService) GetInfo() string {
	return (ps.paymentRepo).GetInfo()
}
