-- name: GetUser :one
SELECT * FROM users
WHERE id = ?;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ?;

-- name: ListUsers :many
SELECT * FROM users
LIMIT ?
OFFSET ?;

-- name: CreateUser :one
INSERT INTO users (
  name, email, hashed_password, bio, auth_token, created_at, updated_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;

