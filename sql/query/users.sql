-- name: CountUsers :one
SELECT COUNT(*) FROM users;

-- name: FindAllUsers :many
SELECT * FROM users
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: FindUserById :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: FindUserByEmail :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;

-- name: FindUserByUsername :one
SELECT * FROM users
WHERE username = $1
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (username, email, password)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET username = $1, email = $2, password = $3
WHERE id = $4
RETURNING *;

-- name: DeleteUser :exec
UPDATE users
SET deleted_at = NOW()
WHERE id = $1;