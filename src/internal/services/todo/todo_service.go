package todo_service

import (
	"github.com/aliaydins/todo/internal/models/todo"
	"github.com/google/uuid"
)

type TodoService struct {
	repository todo.Repository
}

func NewTodoService(repository todo.Repository) *TodoService {
	return &TodoService{repository: repository}
}

func (t *TodoService) Create(req *CreateTodoRequest) error {
	nTodo := todo.Todo{
		ID:   uuid.New(),
		Todo: req.Todo,
	}

	return todo.Create(t.repository, &nTodo)
}

func (t *TodoService) GetList() (*[]todo.Todo, error) {
	return todo.GetList(t.repository)
}

func (t *TodoService) FindById(req *FindByIDTodoRequest) (*todo.Todo, error) {
	return todo.FindById(t.repository, req.ID)
}

func (t *TodoService) Update(req *UpdateTodoRequest) (*todo.Todo, error) {
	return todo.Update(t.repository, req.ID, req.Todo)
}

func (t *TodoService) Approve(req *ApproveTodoRequest) (*todo.Todo, error) {
	return todo.Approve(t.repository, req.ID, req.Done)
}

func (t *TodoService) Delete(req *DeleteTodoRequest) error {
	return todo.Delete(t.repository, req.ID)
}
