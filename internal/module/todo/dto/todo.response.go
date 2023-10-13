package dto

import (
	"github.com/google/uuid"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
)

type Response struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
}

func FromDB(dbTodo db.Todo) Response {
	return Response{
		ID:        dbTodo.ID,
		Title:     dbTodo.Title,
		Completed: dbTodo.Completed,
	}
}

func FromDBList(dbTodos []db.Todo) []Response {
	var todos []Response
	for _, todo := range dbTodos {
		todos = append(todos, FromDB(todo))
	}
	return todos
}
