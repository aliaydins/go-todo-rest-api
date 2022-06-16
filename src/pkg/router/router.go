package router

import (
	"log"

	auth_controller "github.com/aliaydins/todo/internal/controllers/auth"
	todo_controller "github.com/aliaydins/todo/internal/controllers/todo"
	todo_repository "github.com/aliaydins/todo/internal/repositories/todo"
	user_repository "github.com/aliaydins/todo/internal/repositories/user"
	"github.com/aliaydins/todo/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, secretKey string) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.CORS())

	uRepo := user_repository.NewUserRepository(db)
	tRepo := todo_repository.NewTodoRepository(db)

	if uRepo.Migration() != nil {
		log.Fatal("User Migration failed")
	}
	if tRepo.Migration() != nil {
		log.Fatal("Todo Migration failed")
	}

	authC := auth_controller.NewAuthController(secretKey, uRepo)
	todoC := todo_controller.NewTodoService(tRepo)

	auth := r.Group("/")
	{
		auth.GET("/health", authC.Health)
		auth.POST("/login", authC.Login)
		auth.POST("/signup", authC.Signup)
	}

	todo := r.Group("/todo")
	{
		todo.GET("/", middleware.AuthMiddleware(secretKey), todoC.GetList)
		todo.GET("/:id", middleware.AuthMiddleware(secretKey), todoC.FindById)
		todo.POST("/", middleware.AuthMiddleware(secretKey), todoC.CreateTodo)
		todo.PUT("/:id", middleware.AuthMiddleware(secretKey), todoC.Update)
		todo.PUT("/approve/:id", middleware.AuthMiddleware(secretKey), todoC.Approve)
		todo.DELETE("/:id", middleware.AuthMiddleware(secretKey), todoC.Delete)

	}

	return r

}
