
-- name: CreateAccount :one
INSERT INTO accounts (owner, balance, currency)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListAccounts :many
SELECT * FROM accounts
LIMIT $1 OFFSET $2;


-- name: GetAccount :one
SELECT * FROM accounts WHERE id = $1;

-- name: UpdateAccount :one
UPDATE accounts
SET balance = $1
WHERE id = $2
returning *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;
