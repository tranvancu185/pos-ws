package repo

import (
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/model/rq"
	"tranvancu185/vey-pos-ws/internal/uconst"
	"tranvancu185/vey-pos-ws/pkg/utils/utime"
)

type ITableRepo interface {
	GetListTable(params *rq.GetListTableRequest) ([]database.GetListTablesRow, error)
	GetTotalTable(params *rq.GetListTableRequest) (int64, error)
	GetTableByID(id int64) (*database.GetTableByIDRow, error)
	CreateTable(params *rq.CreateTableRequest) (int64, error)
	UpdateTableByID(id int64, params *rq.UpdateTableRequest) error
	DeleteTableByID(id int64) error
}

type tableRepo struct {
	sqlc *database.Queries
}

func NewTableRepo() ITableRepo {
	return &tableRepo{
		sqlc: database.New(global.Mdbc),
	}
}

func (ur *tableRepo) GetListTable(params *rq.GetListTableRequest) ([]database.GetListTablesRow, error) {

	var input database.GetListTablesParams

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
		input.TableCode = params.Text
		input.TableName = params.Text
	}

	if params.TableStatus != 0 {
		input.TableStatus = params.TableStatus
	}

	result, err := ur.sqlc.GetListTables(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ur *tableRepo) GetTotalTable(params *rq.GetListTableRequest) (int64, error) {
	var input database.GetTotalTablesParams

	if params.Text != "" {
		input.TableCode = params.Text
		input.TableName = params.Text
	}

	if params.TableStatus != 0 {
		input.TableStatus = params.TableStatus
	}

	total, err := ur.sqlc.GetTotalTables(ctx, input)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (ur *tableRepo) GetTableByID(id int64) (*database.GetTableByIDRow, error) {
	result, err := ur.sqlc.GetTableByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ur *tableRepo) CreateTable(params *rq.CreateTableRequest) (int64, error) {
	var input database.CreateTableParams
	currentTime := utime.GetCurrentTimeUnix()

	input.TableName = params.TableName
	input.CreatedAt.Int64 = currentTime
	input.UpdatedAt.Int64 = currentTime

	id, err := ur.sqlc.CreateTable(ctx, input)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ur *tableRepo) UpdateTableByID(id int64, params *rq.UpdateTableRequest) error {
	var input database.UpdateTableByIDParams
	currentTime := utime.GetCurrentTimeUnix()

	input.TableID = id

	if params.TableName != "" {
		input.TableName = params.TableName
	}

	if params.TableStatus != 0 {
		input.TableStatus = params.TableStatus
	}

	input.UpdatedAt.Int64 = currentTime

	err := ur.sqlc.UpdateTableByID(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (ur *tableRepo) DeleteTableByID(id int64) error {
	currentTime := utime.GetCurrentTimeUnix()
	var input database.DeleteTableByIDParams
	input.TableID = id
	input.UpdatedAt.Int64 = currentTime

	err := ur.sqlc.DeleteTableByID(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
