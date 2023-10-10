package todo

import "github.com/EdwardKerckhof/gohtmx/internal/domain"

type Todo struct {
	domain.BaseEntity
	Title     string `json:"title" db:"title"`
	Completed bool   `json:"completed" db:"completed"`
}
