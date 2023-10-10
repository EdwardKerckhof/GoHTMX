package ports

import (
	"github.com/google/uuid"

	"github.com/EdwardKerckhof/gohtmx/internal/domain"
)

type TodoStore interface {
	FindAll() ([]domain.Todo, error)
	FindById(id uuid.UUID) (*domain.Todo, error)
	Create(todo *domain.Todo) error
	Update(todo *domain.Todo) error
	Delete(id uuid.UUID) error
}
