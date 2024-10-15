package repo

import (
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
)

type ITableRepo interface {
	GetListTable(params database.GetListTablesParams) ([]database.GetListTablesRow, error)
	GetTotalTable(params database.GetTotalTablesParams) (int64, error)
	GetTableByID(id int64) (*database.GetTableByIDRow, error)
	CreateTable(params database.CreateTableParams) error
	UpdateTableByID(params database.UpdateTableByIDParams) error
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

func (ur *tableRepo) CreateTable(params database.CreateTableParams) error {
	err := ur.sqlc.CreateTable(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (ur *tableRepo) UpdateTableByID(params database.UpdateTableByIDParams) error {
	err := ur.sqlc.UpdateTableByID(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (ur *tableRepo) DeleteTableByID(params database.DeleteTableByIDParams) error {
	err := ur.sqlc.DeleteTableByID(ctx, params)
	if err != nil {
		return err
	}
	return nil
}
