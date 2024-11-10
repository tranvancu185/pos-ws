package service

import (
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/model/rq"
	"tranvancu185/vey-pos-ws/internal/repo"
)

type ICustomerService interface {
	GetListCustomer(params *rq.GetListCustomerRequest) ([]database.GetListCustomersRow, error)
	GetTotalCustomer(params *rq.GetListCustomerRequest) (int64, error)
	GetCustomerByID(id int64) (*database.GetCustomerByIDRow, error)
	CreateCustomer(params *rq.CreateCustomerRequest) (int64, error)
	UpdateCustomerByID(id int64, params *rq.UpdateCustomerRequest) error
	DeleteCustomerByID(id int64) error
}

type customerService struct {
	customerRepo repo.ICustomerRepo
}

func NewCustomerService(customerRepo repo.ICustomerRepo) ICustomerService {
	return &customerService{
		customerRepo: customerRepo,
	}
}

func (cs *customerService) GetListCustomer(params *rq.GetListCustomerRequest) ([]database.GetListCustomersRow, error) {
	customers, err := cs.customerRepo.GetListCustomer(params)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (cs *customerService) CreateCustomer(params *rq.CreateCustomerRequest) (int64, error) {
	code, errCode := NewCommonService().GenerateCode(COUNTER_CUSTOMER)
	if errCode != nil {
		return 0, errCode
	}

	params.CustomerCode = code

	id, err := cs.customerRepo.CreateCustomer(params)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (cs *customerService) UpdateCustomerByID(id int64, params *rq.UpdateCustomerRequest) error {
	err := cs.customerRepo.UpdateCustomerByID(id, params)
	if err != nil {
		return err
	}
	return nil
}

func (cs *customerService) GetCustomerByID(id int64) (*database.GetCustomerByIDRow, error) {
	customer, err := cs.customerRepo.GetCustomerByID(id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (cs *customerService) DeleteCustomerByID(id int64) error {
	err := cs.customerRepo.DeleteCustomerByID(id)
	if err != nil {
		return err
	}
	return nil
}

func (cs *customerService) GetTotalCustomer(params *rq.GetListCustomerRequest) (int64, error) {
	total, err := cs.customerRepo.GetTotalCustomer(params)
	if err != nil {
		return 0, err
	}
	return total, nil
}
