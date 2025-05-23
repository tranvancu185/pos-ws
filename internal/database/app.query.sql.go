// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: app.query.sql

package database

import (
	"context"
	"database/sql"
)

const createApp = `-- name: CreateApp :one
INSERT INTO apps (app_name, app_company, app_version, app_status, app_data, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING app_id
`

type CreateAppParams struct {
	AppName    string         `json:"app_name"`
	AppCompany string         `json:"app_company"`
	AppVersion string         `json:"app_version"`
	AppStatus  sql.NullInt64  `json:"app_status"`
	AppData    sql.NullString `json:"app_data"`
	CreatedAt  sql.NullInt64  `json:"created_at"`
	UpdatedAt  sql.NullInt64  `json:"updated_at"`
}

func (q *Queries) CreateApp(ctx context.Context, arg CreateAppParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createApp,
		arg.AppName,
		arg.AppCompany,
		arg.AppVersion,
		arg.AppStatus,
		arg.AppData,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var app_id int64
	err := row.Scan(&app_id)
	return app_id, err
}

const getAppByID = `-- name: GetAppByID :one
SELECT app_id, app_name, app_company, app_version, app_status, app_data, created_at, updated_at, deleted_at FROM apps WHERE app_id = ?
`

func (q *Queries) GetAppByID(ctx context.Context, appID int64) (App, error) {
	row := q.db.QueryRowContext(ctx, getAppByID, appID)
	var i App
	err := row.Scan(
		&i.AppID,
		&i.AppName,
		&i.AppCompany,
		&i.AppVersion,
		&i.AppStatus,
		&i.AppData,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getListApps = `-- name: GetListApps :many
SELECT app_id, app_name, app_company, app_version, app_status, app_data, created_at, updated_at, deleted_at FROM apps WHERE app_name LIKE ? OR app_company LIKE ? OR app_version = ? OR app_status = ? OR (created_at >= ? AND created_at <= ?) OR (deleted_at >= ? AND deleted_at <= ?) ORDER BY ? LIMIT ? OFFSET ?
`

type GetListAppsParams struct {
	AppName     string        `json:"app_name"`
	AppCompany  string        `json:"app_company"`
	AppVersion  string        `json:"app_version"`
	AppStatus   sql.NullInt64 `json:"app_status"`
	CreatedAt   sql.NullInt64 `json:"created_at"`
	CreatedAt_2 sql.NullInt64 `json:"created_at_2"`
	DeletedAt   sql.NullInt64 `json:"deleted_at"`
	DeletedAt_2 sql.NullInt64 `json:"deleted_at_2"`
	Limit       int64         `json:"limit"`
	Offset      int64         `json:"offset"`
}

func (q *Queries) GetListApps(ctx context.Context, arg GetListAppsParams) ([]App, error) {
	rows, err := q.db.QueryContext(ctx, getListApps,
		arg.AppName,
		arg.AppCompany,
		arg.AppVersion,
		arg.AppStatus,
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
	var items []App
	for rows.Next() {
		var i App
		if err := rows.Scan(
			&i.AppID,
			&i.AppName,
			&i.AppCompany,
			&i.AppVersion,
			&i.AppStatus,
			&i.AppData,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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

const updateAppByID = `-- name: UpdateAppByID :exec
UPDATE apps SET app_name = ?, app_company = ?, app_version = ?, app_status=?, app_data = ?, updated_at = ? WHERE app_id = ?
`

type UpdateAppByIDParams struct {
	AppName    string         `json:"app_name"`
	AppCompany string         `json:"app_company"`
	AppVersion string         `json:"app_version"`
	AppStatus  sql.NullInt64  `json:"app_status"`
	AppData    sql.NullString `json:"app_data"`
	UpdatedAt  sql.NullInt64  `json:"updated_at"`
	AppID      int64          `json:"app_id"`
}

func (q *Queries) UpdateAppByID(ctx context.Context, arg UpdateAppByIDParams) error {
	_, err := q.db.ExecContext(ctx, updateAppByID,
		arg.AppName,
		arg.AppCompany,
		arg.AppVersion,
		arg.AppStatus,
		arg.AppData,
		arg.UpdatedAt,
		arg.AppID,
	)
	return err
}
