-- name: CreateTable :exec
INSERT INTO tables (table_name, table_code, table_status, table_properties, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateTableStatusByID :exec
UPDATE tables
SET table_status = ?,
    updated_at = ?
WHERE table_id = ?;

-- name: UpdateTableByID :exec
UPDATE tables
SET table_name = ?,
    table_code = ?,
    table_status = ?,
    table_properties = ?,
    updated_at = ?
WHERE table_id = ?;

-- name: DeleteTableByID :exec
UPDATE tables
SET deleted_at = ?,
    table_status = 2,
    updated_at = ?
WHERE table_id = ?;

-- name: RestoreTableByID :exec
UPDATE tables
SET deleted_at = 0,
    table_status = 0,
    updated_at = ?
WHERE table_id = ?;

-- name: ForceDeleteTableByID :exec
DELETE FROM tables
WHERE table_id = ? OR table_code = ?;

-- name: GetListTables :many
SELECT table_id, table_name, table_code, table_status, table_properties
FROM tables
WHERE table_name LIKE ? OR table_code like ? OR table_status = ? OR (created_at >= ? AND created_at <= ?) OR (deleted_at >= ? AND deleted_at <= ?)
ORDER BY table_id DESC
LIMIT ? OFFSET ?;

-- name: GetTableByID :one
SELECT table_id, table_name, table_code, table_status, table_properties
FROM tables
WHERE table_id = ?;

-- name: GetTotalTables :one
SELECT COUNT(table_id)
FROM tables
WHERE table_name LIKE ? OR table_code like ? OR table_status = ? OR (created_at >= ? AND created_at <= ?) OR (deleted_at >= ? AND deleted_at <= ?);