-- name: GetListUsers :many
SELECT user_id, user_display_name, user_phone, user_status, user_avatar, user_role_id, created_at, updated_at, deleted_at FROM users;

-- name: GetListUserDeleted :many
SELECT user_id, user_display_name, user_phone, user_status, user_avatar, user_role_id, created_at, updated_at, deleted_at FROM users WHERE user_status = 2;

-- name: GetListUserByFilter :many
SELECT user_id, user_display_name, user_phone, user_avatar, user_status, user_role_id, created_at, updated_at, deleted_at FROM users 
WHERE user_display_name LIKE ? 
OR user_name LIKE ?
OR user_phone LIKE ? 
OR user_status IN (sqlc.slice('user_status_ids')) 
OR user_role_id IN (sqlc.slice('user_role_ids')) 
OR (created_at >= ? AND created_at <= ?) 
OR (deleted_at >= ? AND deleted_at <= ?)
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetTotalUserByFilter :one
SELECT COUNT(user_id) FROM users 
WHERE user_display_name LIKE ?
OR user_phone = ?
OR user_status = ?
OR (created_at >= ? AND created_at <= ?)
OR (deleted_at >= ? AND deleted_at <= ?);

-- name: GetUserByID :one
SELECT user_id, user_display_name,  user_phone, user_avatar, user_status, user_role_id, created_at, updated_at, deleted_at FROM users WHERE user_id = ? LIMIT 1;

-- name: GetUserByPhone :one
SELECT user_id, user_display_name,  user_phone, user_avatar, user_status, user_role_id, created_at, updated_at, deleted_at FROM users WHERE user_phone = ? LIMIT 1;

-- name: GetUserProfile :one
SELECT user_id, user_display_name, user_name, user_phone, user_password, user_status, user_role_id, created_at, updated_at, deleted_at FROM users WHERE user_id = ? OR user_name = ? OR user_Phone = ? LIMIT 1;

-- name: CheckUserExist :one
SELECT COUNT(user_id) FROM users WHERE user_name = ? OR user_phone = ?;

-- name: UpdateUserStatusByID :exec
UPDATE users
SET user_status = ?,
    updated_at = ?
WHERE user_id = ?;

-- name: UpdateUserByID :exec
UPDATE users
SET user_display_name = ?,
    user_phone = ?,
    user_status = ?,
    user_role_id = ?, 
    updated_at = ?
WHERE user_id = ?;

-- name: UpdateUserPasswordByID :exec
UPDATE users
SET user_password = ?,
    updated_at = ?
WHERE user_id = ?;

-- name: UpdateUserAvatarByID :exec
UPDATE users
SET user_avatar = ?,
    updated_at = ?
WHERE user_id = ?;

-- name: CreateUser :one
INSERT INTO users (user_display_name, user_name, user_phone, user_password, user_status, user_avatar, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
RETURNING user_id;

-- name: DeleteUserByID :exec
UPDATE users
SET deleted_at = ?,
    user_status = 2,
    updated_at = ?
WHERE user_id = ?;

-- name: DeleteUserByPhone :exec
UPDATE users
SET deleted_at = ?,
    user_status = 2,
    updated_at = ?
WHERE user_phone = ?;

-- name: RestoreUserByID :exec
UPDATE users
SET deleted_at = 0,
    user_status = 0,
    updated_at = ?
WHERE user_id = ?;

-- name: RestoreUserByPhone :exec
UPDATE users
SET deleted_at = 0,
    user_status = 0,
    updated_at = ?
WHERE user_phone = ?;

-- name: ForceDeleteUserByID :exec
DELETE FROM users WHERE user_id = ?;