package types

import "time"

// here func that interacte with the db
type TodoStore interface {
	GetTodos()
	GetTodoByID()
	UpdateTodo()
	DeleteTodo()
}

type Todo struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string    `json:"task,omitempty" bson:"task,omitempty"`
	Completed bool      `json:"completed" bson:"completed"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt  time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type Models struct {
	Todo Todo
}
