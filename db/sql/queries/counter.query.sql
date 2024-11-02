-- name: CreateCounter :exec
INSERT INTO counters (counter_name, counter_number, created_at, updated_at)
VALUES (?, 1, ?, ?) RETURNING counter_id;

-- name: UpdateCounter :exec
UPDATE counters
SET counter_number = ?,
    updated_at = ?
WHERE counter_name = ?;

-- name: GetCounter :one
SELECT counter_number, updated_at
FROM counters
WHERE counter_name = ?;