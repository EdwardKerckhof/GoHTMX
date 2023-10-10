package domain

type Todo struct {
	BaseEntity
	Title     string `json:"title" db:"title"`
	Completed bool   `json:"completed" db:"completed"`
}
