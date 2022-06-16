package auth_controller

import (
	"net/http"

	"github.com/aliaydins/todo/internal/models/user"
	auth_service "github.com/aliaydins/todo/internal/services/auth"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	secretKey string
	service   auth_service.AuthService
}

func NewAuthController(secretKey string, repository user.Repository) *AuthController {
	return &AuthController{
		secretKey: secretKey,
		service:   *auth_service.NewAuthService(repository),
	}
}

func (c *AuthController) Health(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"message": "go-todo-rest-api service is up"})
}

func (c *AuthController) Login(g *gin.Context) {
	req := new(auth_service.LoginRequest)
	if err := g.BindJSON(req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	accessToken, err := c.service.Authenticate(req, c.secretKey)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

func (c *AuthController) Signup(g *gin.Context) {
	req := new(auth_service.SignupRequest)
	if err := g.BindJSON(req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.service.Signup(req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusCreated, gin.H{"message": "User created"})
}
