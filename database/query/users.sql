-- name: GetUser :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id;

-- name: CreateUser :execresult
INSERT INTO users (name, email, password) VALUES (?, ?, ?);

-- name: UpdateUser :exec
UPDATE users 
    SET name = ?, email = ?, password = ?, updated_at = now() 
    WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;

-- name: GetUserByEmail :one
SELECT id, email FROM users WHERE email = ?;