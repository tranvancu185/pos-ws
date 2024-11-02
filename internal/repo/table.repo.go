package repo

import (
	"time"
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
)

type ITableRepo interface {
	GetListTable(params database.GetListTablesParams) ([]database.GetListTablesRow, error)
	GetTotalTable(params database.GetTotalTablesParams) (int64, error)
	GetTableByID(id int64) (*database.GetTableByIDRow, error)
	CreateTable(params database.CreateTableParams) (int64, error)
	UpdateTableByID(params database.UpdateTableByIDParams) error
	DeleteTableByID(params database.DeleteTableByIDParams) error
}

type tableRepo struct {
	sqlc *database.Queries
}

func NewTableRepo() ITableRepo {
	return &tableRepo{
		sqlc: database.New(global.Mdbc),
	}
}

func (ur *tableRepo) GetListTable(params database.GetListTablesParams) ([]database.GetListTablesRow, error) {
	result, err := ur.sqlc.GetListTables(ctx, params)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ur *tableRepo) GetTotalTable(params database.GetTotalTablesParams) (int64, error) {
	total, err := ur.sqlc.GetTotalTables(ctx, params)
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

func (ur *tableRepo) CreateTable(params database.CreateTableParams) (int64, error) {
	currentTime := time.Now().Unix()
	params.CreatedAt.Int64 = currentTime
	params.UpdatedAt.Int64 = currentTime

	id, err := ur.sqlc.CreateTable(ctx, params)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ur *tableRepo) UpdateTableByID(params database.UpdateTableByIDParams) error {
	currentTime := time.Now().Unix()
	params.UpdatedAt.Int64 = currentTime

	err := ur.sqlc.UpdateTableByID(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (ur *tableRepo) DeleteTableByID(params database.DeleteTableByIDParams) error {
	currentTime := time.Now().Unix()
	params.UpdatedAt.Int64 = currentTime

	err := ur.sqlc.DeleteTableByID(ctx, params)
	if err != nil {
		return err
	}
	return nil
}
