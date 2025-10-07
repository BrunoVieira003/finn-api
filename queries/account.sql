-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY name;

-- name: FindAccountById :one
SELECT * FROM accounts WHERE id = $1;

-- name: CreateAccount :one
INSERT INTO accounts (name, amount)
VALUES ($1, 0)
RETURNING *;