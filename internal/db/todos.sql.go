// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: todos.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const countTodos = `-- name: CountTodos :one
SELECT COUNT(*) FROM todos
WHERE user_id = $1
`

func (q *Queries) CountTodos(ctx context.Context, userID uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, countTodos, userID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (title, user_id)
VALUES ($1, $2)
RETURNING id, title, completed, user_id, created_at, updated_at, deleted_at
`

type CreateTodoParams struct {
	Title  string    `json:"title"`
	UserID uuid.UUID `json:"userId"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRow(ctx, createTodo, arg.Title, arg.UserID)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Completed,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1
`

func (q *Queries) DeleteTodo(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteTodo, id)
	return err
}

const findAllTodos = `-- name: FindAllTodos :many
SELECT id, title, completed, user_id, created_at, updated_at, deleted_at FROM todos
WHERE user_id = $1
ORDER BY id DESC
LIMIT $2
OFFSET $3
`

type FindAllTodosParams struct {
	UserID uuid.UUID `json:"userId"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *Queries) FindAllTodos(ctx context.Context, arg FindAllTodosParams) ([]Todo, error) {
	rows, err := q.db.Query(ctx, findAllTodos, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Todo{}
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Completed,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findTodoById = `-- name: FindTodoById :one
SELECT id, title, completed, user_id, created_at, updated_at, deleted_at FROM todos
WHERE id = $1
LIMIT 1
`

func (q *Queries) FindTodoById(ctx context.Context, id uuid.UUID) (Todo, error) {
	row := q.db.QueryRow(ctx, findTodoById, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Completed,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos
SET title = $1, completed = $2
WHERE id = $3
RETURNING id, title, completed, user_id, created_at, updated_at, deleted_at
`

type UpdateTodoParams struct {
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	ID        uuid.UUID `json:"id"`
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRow(ctx, updateTodo, arg.Title, arg.Completed, arg.ID)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Completed,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
