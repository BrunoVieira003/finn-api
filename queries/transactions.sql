-- name: ListTransactions :many
SELECT * FROM transactions ORDER BY date DESC;

-- name: FindTransactionById :one
SELECT * FROM transactions WHERE id = $1;

-- name: CreateTransaction :one
INSERT INTO transactions (amount, type, account_id, date, description)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;