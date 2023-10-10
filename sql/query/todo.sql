-- name: CountTodos :one
SELECT COUNT(*) FROM todos;

-- name: FindAllTodos :many
SELECT * FROM todos
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: FindTodoById :one
SELECT * FROM todos
WHERE id = $1
LIMIT 1;

-- name: CreateTodo :one
INSERT INTO todos (id, title)
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