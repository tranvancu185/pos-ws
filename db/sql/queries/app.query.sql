-- name: CreateApp :one
INSERT INTO apps (app_name, app_company, app_version, app_status, app_data, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING app_id;

-- name: UpdateAppByID :exec
UPDATE apps SET app_name = ?, app_company = ?, app_version = ?, app_status=?, app_data = ?, updated_at = ? WHERE app_id = ?;

-- name: GetAppByID :one
SELECT app_id, app_name, app_company, app_version, app_status, app_data, created_at, updated_at, deleted_at FROM apps WHERE app_id = ?;

-- name: GetListApps :many
SELECT app_id, app_name, app_company, app_version, app_status, app_data, created_at, updated_at, deleted_at FROM apps WHERE app_name LIKE ? OR app_company LIKE ? OR app_version = ? OR app_status = ? OR (created_at >= ? AND created_at <= ?) OR (deleted_at >= ? AND deleted_at <= ?) ORDER BY ? LIMIT ? OFFSET ?;