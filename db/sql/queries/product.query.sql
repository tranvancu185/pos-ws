-- name: GetListProducts :many
SELECT product_id, product_code, product_name, product_description, product_price, product_status, product_properties 
FROM products 
WHERE product_id = ? 
    AND product_code = ? 
    AND product_status = ? 
    AND (created_at >= ? AND created_at <= ?) 
    AND (deleted_at >= ? AND deleted_at <= ?)
ORDER BY ?
LIMIT ? 
OFFSET ?;

-- name: GetProductByID :one
SELECT product_id, product_name, product_code, product_description, product_price, product_status, product_properties FROM products WHERE product_id = ?;

-- name: GetProductByCode :one
SELECT product_id, product_name, product_code, product_description, product_price, product_status, product_properties FROM products WHERE product_code = ?;

-- name: GetProductByCategoryID :many
SELECT product_id, product_name, product_code, product_description, product_price, product_status, product_properties FROM products WHERE product_category_id = ? ORDER BY ? LIMIT ?;

-- name: SearchProducts :many
SELECT product_id, product_name, product_code, product_description, product_price, product_status, product_properties 
FROM products 
WHERE product_status = 1 
    AND (product_id = ? OR product_name LIKE ? OR product_code LIKE ?)
ORDER BY ? 
LIMIT ? 
OFFSET ?; 

-- name: CreateProduct :exec
INSERT INTO products (product_name, product_code, product_description, product_price, product_status, product_properties) VALUES (?, ?, ?, ?, ?, ?);

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
    updated_at = ?
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