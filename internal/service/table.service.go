package service

import (
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
	tables, err := ts.tableRepo.GetListTable(params)
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (ts *tableService) CreateTable(params *rq.CreateTableRequest) (int64, error) {
	code, errCode := NewCommonService().GenerateCode(COUNTER_TABLE)
	if errCode != nil {
		return 0, errCode
	}

	params.TableCode = code

	id, err := ts.tableRepo.CreateTable(params)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ts *tableService) UpdateTable(id int64, params *rq.UpdateTableRequest) error {
	err := ts.tableRepo.UpdateTableByID(params)
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
	total, err := ts.tableRepo.GetTotalTable(params)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (ts *tableService) DeleteTableByID(id int64) error {
	err := ts.tableRepo.DeleteTableByID(id)
	if err != nil {
		return err
	}
	return nil
}
