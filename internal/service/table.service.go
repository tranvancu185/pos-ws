package service

import (
	"fmt"
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/model/rq"
	"tranvancu185/vey-pos-ws/internal/repo"
)

type ITableService interface {
	GetListTable(params *rq.GetListTableRequest) ([]database.GetListTablesRow, error)
	CreateTable(params *rq.CreateTableRequest) (int64, error)
	UpdateTable(id int64, params *rq.UpdateTableRequest) error
	GetTableByID(id int64) (*database.GetTableByIDRow, error)
	GetTotalTable(params *rq.GetListTableRequest) (int64, error)
	DeleteTableByID(id int64) error
}

type tableService struct {
	tableRepo repo.ITableRepo
}

func NewTableService(tableRepo repo.ITableRepo) ITableService {
	return &tableService{
		tableRepo: tableRepo,
	}
}

func (ts *tableService) GetListTable(params *rq.GetListTableRequest) ([]database.GetListTablesRow, error) {
	var input database.GetListTablesParams

	if params.PageSize != 0 {
		input.Limit = params.PageSize
	} else {
		input.Limit = 10
	}

	if params.Page != 0 {
		input.Offset = (params.Page - 1) * params.PageSize
	} else {
		input.Offset = 0
	}

	if params.Text != "" {
		input.TableCode = params.Text
		input.TableName = params.Text
	}

	if params.TableStatus != 0 {
		input.TableStatus = params.TableStatus
	}

	tables, err := ts.tableRepo.GetListTable(input)
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (ts *tableService) CreateTable(params *rq.CreateTableRequest) (int64, error) {
	var input database.CreateTableParams
	input.TableName = params.TableName

	fmt.Println("input.TableName", input.TableName)
	code, errCode := NewCommonService().GenerateCode(COUNTER_TABLE)
	if errCode != nil {
		return 0, errCode
	}

	fmt.Println("code", code)
	input.TableCode = code

	id, err := ts.tableRepo.CreateTable(input)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ts *tableService) UpdateTable(id int64, params *rq.UpdateTableRequest) error {
	var input database.UpdateTableByIDParams

	if params.TableName != "" {
		input.TableName = params.TableName
	}

	if params.TableStatus != 0 {
		input.TableStatus = params.TableStatus
	}

	err := ts.tableRepo.UpdateTableByID(input)
	if err != nil {
		return err
	}
	return nil
}

func (ts *tableService) GetTableByID(id int64) (*database.GetTableByIDRow, error) {
	table, err := ts.tableRepo.GetTableByID(id)
	if err != nil {
		return nil, err
	}
	return table, nil
}

func (ts *tableService) GetTotalTable(params *rq.GetListTableRequest) (int64, error) {
	var input database.GetTotalTablesParams

	if params.Text != "" {
		input.TableCode = params.Text
		input.TableName = params.Text
	}

	if params.TableStatus != 0 {
		input.TableStatus = params.TableStatus
	}

	total, err := ts.tableRepo.GetTotalTable(input)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (ts *tableService) DeleteTableByID(id int64) error {
	var input database.DeleteTableByIDParams
	input.TableID = id

	err := ts.tableRepo.DeleteTableByID(input)
	if err != nil {
		return err
	}
	return nil
}
