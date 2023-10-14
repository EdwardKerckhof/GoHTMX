-- name: CountTodos :one
SELECT COUNT(*) FROM todos
WHERE user_id = $1;

-- name: FindAllTodos :many
SELECT * FROM todos
WHERE user_id = $1
ORDER BY id DESC
LIMIT $2
OFFSET $3;

-- name: FindTodoById :one
SELECT * FROM todos
WHERE id = $1
LIMIT 1;

-- name: CreateTodo :one
INSERT INTO todos (title, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos
SET title = $1, completed = $2
WHERE id = $3
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;