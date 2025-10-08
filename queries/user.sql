-- name: CreateUser :one
INSERT INTO users (username, email, password)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListUsers :many
SELECT id, username, email FROM users;

-- name: FindUserById :one
SELECT id, username, email FROM users WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;