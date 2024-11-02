package service

import (
	"time"
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/repo"
	"tranvancu185/vey-pos-ws/pkg/request"
)

type ITableService interface {
	GetListTable(params *request.GetListTableRequest) ([]database.GetListTablesRow, error)
}

type tableService struct {
	tableRepo repo.ITableRepo
}

func NewTableService(tableRepo repo.ITableRepo) ITableService {
	return &tableService{
		tableRepo: tableRepo,
	}
}

func (ts *tableService) GetListTable(params *request.GetListTableRequest) ([]database.GetListTablesRow, error) {
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

func (ts *tableService) CreateTable(params *request.CreateTableRequest) error {
	var input database.CreateTableParams

	input.TableCode = params.TableCode
	input.TableName = params.TableName

	err := ts.tableRepo.CreateTable(input)
	if err != nil {
		return err
	}
	return nil
}

func (ts *tableService) UpdateTable(params *request.UpdateTableRequest) error {
	var input database.UpdateTableByIDParams
	input.TableID = params.TableID

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

func (ts *tableService) GetTotalTable(params *request.GetListTableRequest) (int64, error) {
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

	currentTimeStamp := time.Now().Unix()
	input.TableID = id
	input.DeletedAt.Int64 = currentTimeStamp

	err := ts.tableRepo.DeleteTableByID(input)
	if err != nil {
		return err
	}
	return nil
}
