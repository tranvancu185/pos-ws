-- name: GetListProducts :many
SELECT product_id, product_code, product_name, product_description, product_price, product_status, product_properties, created_at, updated_at, deleted_at, product_display_name, product_image, product_category_id
FROM products 
WHERE
    (
        product_code like ? 
        OR product_name like ?
    )
    AND product_status = ? 
    AND (created_at >= ? AND created_at <= ?) 
    AND (deleted_at >= ? AND deleted_at <= ?)
    AND product_category_id = ?
ORDER BY ?
LIMIT ? 
OFFSET ?;

-- name: GetTotalProducts :one
SELECT COUNT(product_id)
FROM products
WHERE product_id = ? 
AND product_code like ?
OR product_name like ?
AND product_status = ? 
AND (created_at >= ? AND created_at <= ?) 
AND (deleted_at >= ? AND deleted_at <= ?)
AND product_category_id = ?;

-- name: GetProductByID :one
SELECT product_id, product_code, product_name, product_description, product_price, product_status, product_properties, created_at, updated_at, deleted_at, product_display_name, product_image, product_category_id FROM products WHERE product_id = ?;

-- name: GetProductByCode :one
SELECT product_id, product_code, product_name, product_description, product_price, product_status, product_properties, created_at, updated_at, deleted_at, product_display_name, product_image, product_category_id FROM products WHERE product_code = ?;

-- name: GetProductByCategoryID :many
SELECT product_id, product_code, product_name, product_description, product_price, product_status, product_properties, created_at, updated_at, deleted_at, product_display_name, product_image, product_category_id FROM products WHERE product_category_id = ? ORDER BY ? LIMIT ?;

-- name: CreateProduct :one
INSERT INTO products (product_name, product_code, product_display_name, product_description, product_price, product_status, product_properties, product_category_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING product_id;

-- name: UpdateProductStatusByID :exec
UPDATE products
SET product_status = ?,
    updated_at = ?
WHERE product_id = ?;

-- name: UpdateProductByID :exec
UPDATE products
SET product_name = ?,
    product_code = ?,
    product_description = ?,
    product_price = ?,
    product_status = ?,
    product_properties = ?,
    updated_at = ?,
    product_category_id = ?,
    product_display_name = ?
WHERE product_id = ?;

-- name: DeleteProductByID :exec
UPDATE products
SET deleted_at = ?,
    product_status = 2,
    updated_at = ?
WHERE product_id = ?;

-- name: RestoreProductByID :exec
UPDATE products
SET deleted_at = 0,
    product_status = 0,
    updated_at = ?
WHERE product_id = ?;

-- name: ForceDeleteProductByID :exec
DELETE FROM products
WHERE product_id = ? OR product_code = ?;

-- name: SearchProducts :many
SELECT product_id, product_name, product_code, product_description, product_price, product_status, product_properties
FROM products
WHERE product_status = ?
AND 
(product_name LIKE ?
OR product_code LIKE ?);

-- name: UpdateProductImageByID :exec
UPDATE products
SET product_image = ?,
    updated_at = ?
WHERE product_id = ?;
