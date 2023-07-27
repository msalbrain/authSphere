-- name: GetSAuthor :one
SELECT * FROM users
WHERE id = ?;

-- name: ListSAuthors :many
SELECT * FROM users
ORDER BY name;

-- name: CreateSAuthor :one
INSERT INTO users (
  name, bio
) VALUES (
  ?, ?
)
RETURNING *;

-- name: DeleteSAuthor :exec
DELETE FROM users
WHERE id = ?;

-- name: GetSJoin :one
SELECT
  users.id AS user_id,
  users.name,
  users.email,
  users.bio,
  user_sessions.session_token,
  user_sessions.login_time,
  user_sessions.logout_time,
  user_sessions.ip_address,
  user_sessions.user_agent,
  user_sessions.is_active
FROM
  users
JOIN
  user_sessions ON users.id = user_sessions.user_id;
