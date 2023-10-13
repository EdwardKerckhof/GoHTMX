package todo

import (
	"time"

	"github.com/google/uuid"

	"github.com/EdwardKerckhof/gohtmx/internal/db"
)

type Todo struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	Title     string     `json:"title" db:"title"`
	Completed bool       `json:"completed" db:"completed"`
	UserID    uuid.UUID  `json:"userId" db:"user_id"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
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
