-- name: CreateAuthor :one
INSERT INTO account (
  owner,
  balance,currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAuthor :one
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM account
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAuthor :exec
UPDATE account
  set balance = $2
WHERE id = $1;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1;

