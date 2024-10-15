package service

import (
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
