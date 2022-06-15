package router

import (
	"log"

	auth_controller "github.com/aliaydins/todo/internal/controllers/auth"
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

	if uRepo.Migration() != nil {
		log.Fatal("User Migration failed")
	}

	authC := auth_controller.NewAuthController(secretKey, uRepo)

	auth := r.Group("/")
	{
		auth.GET("/health", authC.Health)
		auth.POST("/login", authC.Login)
		auth.POST("/signup", authC.Signup)
	}

	return r

}
