-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY name;

-- name: FindAccountById :one
SELECT
    A.id,
    A.name,
    COALESCE(SUM(CASE WHEN t.type = 'income' THEN t.amount END), 0) AS total_income,
    COALESCE(SUM(CASE WHEN t.type = 'expense' THEN t.amount END), 0) AS total_expense,
    COALESCE(
        SUM(
            CASE
                WHEN T.type = 'income' THEN T.amount
                WHEN T.type = 'expense' THEN -T.amount
                ELSE 0
            END
        ),
        0
    ) AS total
FROM accounts A
LEFT JOIN transactions T ON T.account_id = A.id
WHERE A.id = $1
GROUP BY A.id, A.name;

-- name: CreateAccount :one
INSERT INTO accounts (name)
VALUES ($1)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;