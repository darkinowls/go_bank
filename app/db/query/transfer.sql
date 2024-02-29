-- name: CreateTransfer :one
INSERT INTO transfers (from_account_id, to_account_id, amount)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListTransfers :many
SELECT *
FROM transfers
LIMIT $1 OFFSET $2;


-- name: GetTransfer :one
SELECT *
FROM transfers
WHERE id = $1;
