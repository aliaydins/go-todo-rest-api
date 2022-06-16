package todo

import "github.com/google/uuid"

type Repository interface {
	Migration() error
	GetList() (*[]Todo, error)
	Create(todo *Todo) error
	FindById(id uuid.UUID) (*Todo, error)
	Update(id uuid.UUID, todo string) (*Todo, error)
	Approve(id uuid.UUID, done bool) (*Todo, error)
	Delete(id uuid.UUID) error
}
