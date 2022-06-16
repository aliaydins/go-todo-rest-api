package todo_controller

import (
	"fmt"
	"net/http"

	"github.com/aliaydins/todo/internal/models/todo"
	todo_service "github.com/aliaydins/todo/internal/services/todo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoController struct {
	service todo_service.TodoService
}

func NewTodoService(r todo.Repository) *TodoController {
	return &TodoController{
		service: *todo_service.NewTodoService(r),
	}
}

func (t *TodoController) CreateTodo(g *gin.Context) {
	req := new(todo_service.CreateTodoRequest)
	if err := g.Bind(req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := t.service.Create(req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusCreated, gin.H{"message": "Todo created"})
}

func (t *TodoController) GetList(g *gin.Context) {
	todos, err := t.service.GetList()
	if err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
		return
	}
	g.JSON(http.StatusOK, todos)
}

func (t *TodoController) FindById(g *gin.Context) {
	req := new(todo_service.FindByIDTodoRequest)
	req.ID = uuid.Must(uuid.Parse(g.Param("id")))

	todo, err := t.service.FindById(req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, todo)

}

func (t *TodoController) Update(g *gin.Context) {
	req := new(todo_service.UpdateTodoRequest)
	req.ID = uuid.Must(uuid.Parse(g.Param("id")))

	if err := g.BindJSON(req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := t.service.Update(req)
	if err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
		return
	}
	g.JSON(http.StatusOK, todo)
}

func (t *TodoController) Approve(g *gin.Context) {
	req := new(todo_service.ApproveTodoRequest)
	req.ID = uuid.MustParse(g.Param("id"))

	if err := g.BindJSON(req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := t.service.Approve(req)

	if err != nil {
		fmt.Println(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, todo)

}

func (t *TodoController) Delete(g *gin.Context) {
	req := new(todo_service.DeleteTodoRequest)
	req.ID = uuid.Must(uuid.Parse(g.Param("id")))

	err := t.service.Delete(req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
