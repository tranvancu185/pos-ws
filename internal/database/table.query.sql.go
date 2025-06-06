// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: table.query.sql

package database

import (
	"context"
	"database/sql"
)

const createTable = `-- name: CreateTable :one
INSERT INTO tables (table_name, table_code, table_status, table_properties, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?, ?) RETURNING table_id
`

type CreateTableParams struct {
	TableName       string         `json:"table_name"`
	TableCode       string         `json:"table_code"`
	TableStatus     int64          `json:"table_status"`
	TableProperties sql.NullString `json:"table_properties"`
	CreatedAt       sql.NullInt64  `json:"created_at"`
	UpdatedAt       sql.NullInt64  `json:"updated_at"`
}

func (q *Queries) CreateTable(ctx context.Context, arg CreateTableParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createTable,
		arg.TableName,
		arg.TableCode,
		arg.TableStatus,
		arg.TableProperties,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var table_id int64
	err := row.Scan(&table_id)
	return table_id, err
}

const deleteTableByID = `-- name: DeleteTableByID :exec
UPDATE tables
SET deleted_at = ?,
    table_status = 2,
    updated_at = ?
WHERE table_id = ?
`

type DeleteTableByIDParams struct {
	DeletedAt sql.NullInt64 `json:"deleted_at"`
	UpdatedAt sql.NullInt64 `json:"updated_at"`
	TableID   int64         `json:"table_id"`
}

func (q *Queries) DeleteTableByID(ctx context.Context, arg DeleteTableByIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteTableByID, arg.DeletedAt, arg.UpdatedAt, arg.TableID)
	return err
}

const forceDeleteTableByID = `-- name: ForceDeleteTableByID :exec
DELETE FROM tables
WHERE table_id = ? OR table_code = ?
`

type ForceDeleteTableByIDParams struct {
	TableID   int64  `json:"table_id"`
	TableCode string `json:"table_code"`
}

func (q *Queries) ForceDeleteTableByID(ctx context.Context, arg ForceDeleteTableByIDParams) error {
	_, err := q.db.ExecContext(ctx, forceDeleteTableByID, arg.TableID, arg.TableCode)
	return err
}

const getListTables = `-- name: GetListTables :many
SELECT table_id, table_name, table_code, table_status, table_properties
FROM tables
WHERE (table_name LIKE ? OR table_code like ?) AND table_status = ? AND (created_at >= ? AND created_at <= ?) AND (deleted_at >= ? AND deleted_at <= ?)
ORDER BY table_id DESC
LIMIT ? OFFSET ?
`

type GetListTablesParams struct {
	TableName   string        `json:"table_name"`
	TableCode   string        `json:"table_code"`
	TableStatus int64         `json:"table_status"`
	CreatedAt   sql.NullInt64 `json:"created_at"`
	CreatedAt_2 sql.NullInt64 `json:"created_at_2"`
	DeletedAt   sql.NullInt64 `json:"deleted_at"`
	DeletedAt_2 sql.NullInt64 `json:"deleted_at_2"`
	Limit       int64         `json:"limit"`
	Offset      int64         `json:"offset"`
}

type GetListTablesRow struct {
	TableID         int64          `json:"table_id"`
	TableName       string         `json:"table_name"`
	TableCode       string         `json:"table_code"`
	TableStatus     int64          `json:"table_status"`
	TableProperties sql.NullString `json:"table_properties"`
}

func (q *Queries) GetListTables(ctx context.Context, arg GetListTablesParams) ([]GetListTablesRow, error) {
	rows, err := q.db.QueryContext(ctx, getListTables,
		arg.TableName,
		arg.TableCode,
		arg.TableStatus,
		arg.CreatedAt,
		arg.CreatedAt_2,
		arg.DeletedAt,
		arg.DeletedAt_2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetListTablesRow
	for rows.Next() {
		var i GetListTablesRow
		if err := rows.Scan(
			&i.TableID,
			&i.TableName,
			&i.TableCode,
			&i.TableStatus,
			&i.TableProperties,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTableByID = `-- name: GetTableByID :one
SELECT table_id, table_name, table_code, table_status, table_properties
FROM tables
WHERE table_id = ?
`

type GetTableByIDRow struct {
	TableID         int64          `json:"table_id"`
	TableName       string         `json:"table_name"`
	TableCode       string         `json:"table_code"`
	TableStatus     int64          `json:"table_status"`
	TableProperties sql.NullString `json:"table_properties"`
}

func (q *Queries) GetTableByID(ctx context.Context, tableID int64) (GetTableByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getTableByID, tableID)
	var i GetTableByIDRow
	err := row.Scan(
		&i.TableID,
		&i.TableName,
		&i.TableCode,
		&i.TableStatus,
		&i.TableProperties,
	)
	return i, err
}

const getTotalTables = `-- name: GetTotalTables :one
SELECT COUNT(table_id)
FROM tables
WHERE (table_name LIKE ? OR table_code like ?) AND table_status = ? AND (created_at >= ? AND created_at <= ?) AND (deleted_at >= ? AND deleted_at <= ?)
`

type GetTotalTablesParams struct {
	TableName   string        `json:"table_name"`
	TableCode   string        `json:"table_code"`
	TableStatus int64         `json:"table_status"`
	CreatedAt   sql.NullInt64 `json:"created_at"`
	CreatedAt_2 sql.NullInt64 `json:"created_at_2"`
	DeletedAt   sql.NullInt64 `json:"deleted_at"`
	DeletedAt_2 sql.NullInt64 `json:"deleted_at_2"`
}

func (q *Queries) GetTotalTables(ctx context.Context, arg GetTotalTablesParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalTables,
		arg.TableName,
		arg.TableCode,
		arg.TableStatus,
		arg.CreatedAt,
		arg.CreatedAt_2,
		arg.DeletedAt,
		arg.DeletedAt_2,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const restoreTableByID = `-- name: RestoreTableByID :exec
UPDATE tables
SET deleted_at = 0,
    table_status = 0,
    updated_at = ?
WHERE table_id = ?
`

type RestoreTableByIDParams struct {
	UpdatedAt sql.NullInt64 `json:"updated_at"`
	TableID   int64         `json:"table_id"`
}

func (q *Queries) RestoreTableByID(ctx context.Context, arg RestoreTableByIDParams) error {
	_, err := q.db.ExecContext(ctx, restoreTableByID, arg.UpdatedAt, arg.TableID)
	return err
}

const updateTableByID = `-- name: UpdateTableByID :exec
UPDATE tables
SET table_name = ?,
    table_code = ?,
    table_status = ?,
    table_properties = ?,
    updated_at = ?
WHERE table_id = ?
`

type UpdateTableByIDParams struct {
	TableName       string         `json:"table_name"`
	TableCode       string         `json:"table_code"`
	TableStatus     int64          `json:"table_status"`
	TableProperties sql.NullString `json:"table_properties"`
	UpdatedAt       sql.NullInt64  `json:"updated_at"`
	TableID         int64          `json:"table_id"`
}

func (q *Queries) UpdateTableByID(ctx context.Context, arg UpdateTableByIDParams) error {
	_, err := q.db.ExecContext(ctx, updateTableByID,
		arg.TableName,
		arg.TableCode,
		arg.TableStatus,
		arg.TableProperties,
		arg.UpdatedAt,
		arg.TableID,
	)
	return err
}

const updateTableStatusByID = `-- name: UpdateTableStatusByID :one
UPDATE tables
SET table_status = ?,
    updated_at = ?
WHERE table_id = ?
RETURNING table_id
`

type UpdateTableStatusByIDParams struct {
	TableStatus int64         `json:"table_status"`
	UpdatedAt   sql.NullInt64 `json:"updated_at"`
	TableID     int64         `json:"table_id"`
}

func (q *Queries) UpdateTableStatusByID(ctx context.Context, arg UpdateTableStatusByIDParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, updateTableStatusByID, arg.TableStatus, arg.UpdatedAt, arg.TableID)
	var table_id int64
	err := row.Scan(&table_id)
	return table_id, err
}
