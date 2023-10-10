package ports

import (
	"github.com/google/uuid"

	"github.com/EdwardKerckhof/gohtmx/internal/domain/todo"
)

type TodoStore interface {
	FindAll() ([]todo.Todo, error)
	FindById(id uuid.UUID) (*todo.Todo, error)
	Create(todo *todo.Todo) error
	Update(todo *todo.Todo) error
	Delete(id uuid.UUID) error
}
