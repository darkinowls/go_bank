
-- name: CreateEntry :one
INSERT INTO entries (account_id, amount)
VALUES ($1, $2)
RETURNING *;

-- name: ListEntries :many
SELECT * FROM entries
LIMIT $1
OFFSET $2;


-- name: GetEntry :one
SELECT * FROM entries WHERE id = $1;
