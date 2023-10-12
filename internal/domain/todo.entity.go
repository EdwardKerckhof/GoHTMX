package domain

import "github.com/google/uuid"

type Todo struct {
	BaseEntity
	UserID    uuid.UUID `json:"userId" db:"user_id"`
	Title     string    `json:"title" db:"title"`
	Completed bool      `json:"completed" db:"completed"`
}
