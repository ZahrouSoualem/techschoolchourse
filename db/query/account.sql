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

-- name: GetAuthorForUpdate :one
SELECT * FROM account
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAuthors :many
SELECT * FROM account
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;


-- name: UpdateAuthor :one
UPDATE account
  set balance = $2
WHERE id = $1
RETURNING *;

-- name: UpdateAuthorBalance :one
UPDATE account
  set balance = balance + sqlc.arg(amount) 
WHERE id = sqlc.arg(id) 
RETURNING *;


-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1;

