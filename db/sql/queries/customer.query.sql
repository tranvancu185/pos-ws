-- name: GetCustomerByID :one
SELECT customer_id, customer_name, customer_email, customer_phone, customer_status, customer_total_orders, customer_properties
FROM customers
WHERE customer_id = ?;

-- name: GetCustomerByPhone :one
SELECT customer_id, customer_name, customer_email, customer_phone, customer_status, customer_total_orders, customer_properties
FROM customers
WHERE customer_phone = ?;

-- name: GetListCustomers :many
SELECT customer_id, customer_name, customer_email, customer_phone, customer_status, customer_total_orders
FROM customers
WHERE customer_name LIKE ? OR customer_phone = ? OR customer_status = ? OR (created_at >= ? AND created_at <= ?) OR (deleted_at >= ? AND deleted_at <= ?)
ORDER BY ?
LIMIT ? OFFSET ?;

-- name: CreateCustomer :exec
INSERT INTO customers (customer_name, customer_email, customer_phone, customer_status, customer_properties) VALUES (?, ?, ?, ?, ?);

-- name: UpdateCustomerStatusByID :exec
UPDATE customers
SET customer_status = ?,
    updated_at = ?
WHERE customer_id = ?;

-- name: UpdateCustomerByID :exec
UPDATE customers
SET customer_name = ?,
    customer_email = ?,
    customer_phone = ?,
    customer_status = ?,
    customer_properties = ?,
    updated_at = ?
WHERE customer_id = ?;

-- name: UpdateCustomerTotalOrdersByID :exec
UPDATE customers
SET customer_total_orders = ?,
    updated_at = ?
WHERE customer_id = ?;

-- name: DeleteCustomerByID :exec
UPDATE customers
SET deleted_at = ?,
    customer_status = 2,
    updated_at = ?
WHERE customer_id = ?;

-- name: RestoreCustomerByID :exec
UPDATE customers
SET deleted_at = 0,
    customer_status = 0,
    updated_at = ?
WHERE customer_id = ?;

-- name: ForceDeleteCustomerByID :exec
DELETE FROM customers
WHERE customer_id = ? OR customer_phone = ?;