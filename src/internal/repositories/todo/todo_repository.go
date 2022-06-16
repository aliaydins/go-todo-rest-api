package todo_repository

import (
	"fmt"

	"github.com/aliaydins/todo/internal/models/todo"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *todoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) Migration() error {
	err := r.db.AutoMigrate(&todo.Todo{})
	if err != nil {
		return err
	}
	return nil
}
func (r *todoRepository) GetList() (*[]todo.Todo, error) {
	var todos []todo.Todo
	err := r.db.Find(&todos).Error
	return &todos, err
}

func (r *todoRepository) Create(todo *todo.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) FindById(id uuid.UUID) (*todo.Todo, error) {
	todo := new(todo.Todo)
	err := r.db.Where("id = ?", id).First(&todo).Error

	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *todoRepository) Update(id uuid.UUID, todo string) (*todo.Todo, error) {

	t, err := r.FindById(id)
	if err != nil {
		return t, err
	}

	err = r.db.Model(&t).Update("todo", todo).Error
	return t, err
}

func (r *todoRepository) Approve(id uuid.UUID, done bool) (*todo.Todo, error) {
	t, err := r.FindById(id)
	if err != nil {
		return t, err
	}
	err = r.db.Model(&t).Update("done", done).Error
	return t, err
}

func (r *todoRepository) Delete(id uuid.UUID) error {
	t := r.db.Where("id = ?", id).Delete(&todo.Todo{})
	fmt.Println(t)

	return t.Error

}
