-- name: CreateUser :execlastid
INSERT INTO users (id, first_name, last_name, email, password)
VALUES (UUID_TO_BIN(UUID(), 1), ?, ?, ?, ?);

-- name: GetUserByID :one
SELECT BIN_TO_UUID(id, 1) AS uuid, * FROM users
WHERE id = UUID_TO_BIN(?, 1)
LIMIT 1;

-- name: GetUserByEmail :one
SELECT BIN_TO_UUID(id, 1) AS uuid, * FROM users
WHERE email = ?
LIMIT 1;

-- name: UpdateUser :exec
UPDATE users
SET
  first_name=coalesce(sqlc.narg('first_name'), first_name),
  last_name=coalesce(sqlc.narg('last_name'), last_name),
  email=coalesce(sqlc.narg('email'), email),
  password=coalesce(sqlc.narg('password'), password)
WHERE id = UUID_TO_BIN(sqlc.arg('id'));

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
