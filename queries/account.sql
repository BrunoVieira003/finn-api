-- name: ListAccounts :many
SELECT id, name, amount FROM accounts
WHERE owner_id = $1
ORDER BY name;

-- name: FindAccountById :one
SELECT id, name, amount FROM accounts
WHERE owner_id = $1 AND id = $2 ;

-- name: CreateAccount :one
INSERT INTO accounts (owner_id, name, amount)
VALUES ($1, $2, 0)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE owner_id = $1 AND id = $2;