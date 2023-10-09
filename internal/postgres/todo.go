package postgres

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/EdwardKerckhof/gohtmx/internal/domain"
	"github.com/EdwardKerckhof/gohtmx/internal/ports"
)

type todoStoreImpl struct {
	*sqlx.DB
}

func NewTodoStore(db *sqlx.DB) ports.TodoStore {
	return &todoStoreImpl{
		db,
	}
}

func (s *todoStoreImpl) FindAll() ([]domain.Todo, error) {
	todos := []domain.Todo{}
	if err := s.Select(&todos, "SELECT * FROM todos WHERE deleted_at IS NULL"); err != nil {
		return todos, fmt.Errorf("error getting todos: %w", err)
	}
	return todos, nil
}

func (s *todoStoreImpl) FindById(id uuid.UUID) (*domain.Todo, error) {
	todo := domain.Todo{}
	if err := s.Get(&todo, "SELECT * FROM todos WHERE id = $1", id); err != nil {
		return &todo, fmt.Errorf("error getting todo: %w", err)
	}
	return &todo, nil
}

func (s *todoStoreImpl) Create(todo *domain.Todo) error {
	if err := s.Get(todo, "INSERT INTO todos (id, title, completed) VALUES ($1, $2, $3) RETURNING *",
		todo.ID,
		todo.Title,
		todo.Completed); err != nil {
		return fmt.Errorf("error creating todo: %w", err)
	}
	return nil
}

func (s *todoStoreImpl) Update(todo *domain.Todo) error {
	if err := s.Get(todo, "UPDATE todos SET title = $1, completed = $2, updated_at = $3 WHERE id = $4 RETURNING *",
		todo.Title,
		todo.Completed,
		time.Now(),
		todo.ID); err != nil {
		return fmt.Errorf("error updating todo: %w", err)
	}
	return nil
}

func (s *todoStoreImpl) Delete(id uuid.UUID) error {
	_, err := s.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting todo: %w", err)
	}
	return nil
}
