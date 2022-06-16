package todo

import (
	"github.com/google/uuid"
)

func Create(r Repository, t *Todo) error {
	return r.Create(t)
}

func GetList(r Repository) (*[]Todo, error) {
	todos, err := r.GetList()
	return todos, err
}

func FindById(r Repository, id uuid.UUID) (*Todo, error) {
	todo, err := r.FindById(id)
	return todo, err
}

func Update(r Repository, id uuid.UUID, todo string) (*Todo, error) {
	t, err := r.Update(id, todo)
	return t, err
}

func Approve(r Repository, id uuid.UUID, done bool) (*Todo, error) {
	t, err := r.Approve(id, done)
	return t, err
}

func Delete(r Repository, id uuid.UUID) error {
	return r.Delete(id)
}
