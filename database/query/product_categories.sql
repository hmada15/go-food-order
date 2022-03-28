-- name: GetProductCategory :one
SELECT * FROM product_categories WHERE id = ? LIMIT 1;

-- name: ListProductCategories :many
SELECT * FROM product_categories ORDER BY created_at;

-- name: CreateProductCategory :execresult
INSERT INTO product_categories 
    (name, description, slug, is_publish) 
    VALUES (?, ?, ?, ?);

-- name: UpdateProductCategory :exec
UPDATE product_categories 
    SET name = ?, description = ?, slug = ?, is_publish = ?, updated_at = now() 
    WHERE id = ?;

-- name: DeleteProductCategory :exec
DELETE FROM product_categories WHERE id = ?;