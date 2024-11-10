package repo

import (
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/model/rq"
	"tranvancu185/vey-pos-ws/internal/uconst"
	"tranvancu185/vey-pos-ws/pkg/utils/utime"
)

type ICustomerRepo interface {
	GetListCustomer(params *rq.GetListCustomerRequest) ([]database.GetListCustomersRow, error)
	GetTotalCustomer(params *rq.GetListCustomerRequest) (int64, error)
	GetCustomerByID(id int64) (*database.GetCustomerByIDRow, error)
	CreateCustomer(params *rq.CreateCustomerRequest) (int64, error)
	UpdateCustomerByID(id int64, params *rq.UpdateCustomerRequest) error
	DeleteCustomerByID(id int64) error
}

type customerRepo struct {
	sqlc *database.Queries
}

func NewCustomerRepo() ICustomerRepo {
	return &customerRepo{
		sqlc: database.New(global.Mdbc),
	}
}

func (ur *customerRepo) GetListCustomer(params *rq.GetListCustomerRequest) ([]database.GetListCustomersRow, error) {

	var input database.GetListCustomersParams

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

	if params.Text != "" {
		// input.CustomerCode = params.Text
		input.CustomerName = params.Text
	}

	if params.CustomerStatus != 0 {
		input.CustomerStatus = params.CustomerStatus
	}

	result, err := ur.sqlc.GetListCustomers(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ur *customerRepo) GetTotalCustomer(params *rq.GetListCustomerRequest) (int64, error) {

	var input database.GetTotalCustomersParams

	if params.Text != "" {
		// input.CustomerCode = params.Text
		input.CustomerName = params.Text
	}

	if params.CustomerStatus != 0 {
		input.CustomerStatus = params.CustomerStatus
	}

	total, err := ur.sqlc.GetTotalCustomers(ctx, input)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (ur *customerRepo) GetCustomerByID(id int64) (*database.GetCustomerByIDRow, error) {
	result, err := ur.sqlc.GetCustomerByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ur *customerRepo) CreateCustomer(params *rq.CreateCustomerRequest) (int64, error) {
	var input database.CreateCustomerParams

	input.CustomerName = params.CustomerName
	input.CustomerPhone = params.CustomerPhone
	input.CustomerEmail.String = params.CustomerEmail
	input.CustomerStatus = uconst.CUSTOMER_STATUS_ACTIVE

	id, err := ur.sqlc.CreateCustomer(ctx, input)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ur *customerRepo) UpdateCustomerByID(id int64, params *rq.UpdateCustomerRequest) error {
	var input database.UpdateCustomerByIDParams

	input.CustomerID = id

	if params.CustomerName != "" {
		input.CustomerName = params.CustomerName
	}

	if params.CustomerEmail != "" {
		input.CustomerEmail.String = params.CustomerEmail
	}

	if params.CustomerStatus != 0 {
		input.CustomerStatus = params.CustomerStatus
	}

	input.UpdatedAt.Int64 = utime.GetCurrentTimeUnix()
	err := ur.sqlc.UpdateCustomerByID(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (ur *customerRepo) DeleteCustomerByID(id int64) error {
	var input database.DeleteCustomerByIDParams
	currentTime := utime.GetCurrentTimeUnix()

	input.CustomerID = id
	input.DeletedAt.Int64 = currentTime

	err := ur.sqlc.DeleteCustomerByID(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
