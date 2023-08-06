-- name: CreateUser :execlastid
INSERT INTO users (first_name, last_name, email, password)
VALUES ($1, $2, $3, $4);

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;

-- name: UpdateUser :exec
UPDATE users
SET
  first_name=coalesce(sqlc.narg('first_name'), first_name),
  last_name=coalesce(sqlc.narg('last_name'), last_name),
  email=coalesce(sqlc.narg('email'), email),
  password=coalesce(sqlc.narg('password'), password)
WHERE id = sqlc.arg('id')::UUID;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
