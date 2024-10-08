-- name: FindSessionById :one
SELECT * FROM sessions
WHERE id = $1
LIMIT 1;

-- name: CreateSession :one
INSERT INTO sessions (id, user_id, refresh_token, user_agent, client_ip, expires_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;