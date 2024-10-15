package service

import "tranvancu185/vey-pos-ws/internal/repo"

type ICustomerService interface {
	GetInfo() string
}

type customerService struct {
	customerRepo repo.ICustomerRepo
}

func NewCustomerService(customerRepo repo.ICustomerRepo) ICustomerService {
	return &customerService{
		customerRepo: customerRepo,
	}
}

// GetInfo implements ICustomer.
func (cus *customerService) GetInfo() string {
	return (cus.customerRepo).GetInfo()
}
