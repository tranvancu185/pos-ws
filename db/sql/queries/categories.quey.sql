-- name: GetListCategories :many
SELECT category_id, category_name, category_description, category_status, category_properties
FROM categories
WHERE category_name LIKE ? OR category_status = ? OR (created_at >= ? AND created_at <= ?) OR (deleted_at >= ? AND deleted_at <= ?)
ORDER BY ?
LIMIT ? OFFSET ?;

-- name: GetCategoryByID :one
SELECT category_id, category_name, category_description, category_status, category_properties
FROM categories
WHERE category_id = ?;

-- name: CreateCategory :exec
INSERT INTO categories (category_name, category_description, category_status, category_properties) VALUES (?, ?, ?, ?);

-- name: UpdateCategoryStatusByID :exec
UPDATE categories
SET category_status = ?,
    updated_at = ?
WHERE category_id = ?;

-- name: UpdateCategoryByID :exec
UPDATE categories
SET category_name = ?,
    category_description = ?,
    category_status = ?,
    category_properties = ?,
    updated_at = ?
WHERE category_id = ?;

-- name: DeleteCategoryByID :exec
UPDATE categories
SET deleted_at = ?,
    category_status = 2,
    updated_at = ?
WHERE category_id = ?;

-- name: RestoreCategoryByID :exec
UPDATE categories
SET deleted_at = 0,
    category_status = 0,
    updated_at = ?
WHERE category_id = ?;

-- name: ForceDeleteCategoryByID :exec
DELETE FROM categories
WHERE category_id = ? OR category_name = ?;