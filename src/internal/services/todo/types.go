package todo_service

import "github.com/google/uuid"

type CreateTodoRequest struct {
	Todo string `json:"todo"`
}

type FindByIDTodoRequest struct {
	ID uuid.UUID `json:"id"`
}

type UpdateTodoRequest struct {
	ID   uuid.UUID `json:"id"`
	Todo string    `json:"todo"`
}

type DeleteTodoRequest struct {
	ID uuid.UUID `json:"id"`
}

type ApproveTodoRequest struct {
	ID   uuid.UUID `json:"id"`
	Done bool      `json:"done"`
}
