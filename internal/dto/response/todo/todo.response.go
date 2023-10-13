package todo

import (
	"github.com/google/uuid"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
)

type Todo struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
}

func FromDB(todo db.Todo) Todo {
	return Todo{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
	}
}

func FromDBList(todos []db.Todo) []Todo {
	var todoDTOs []Todo
	for _, todo := range todos {
		todoDTOs = append(todoDTOs, FromDB(todo))
	}
	return todoDTOs
}
