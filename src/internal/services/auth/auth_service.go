package auth_service

import (
	"github.com/aliaydins/todo/internal/models/user"
)

type AuthService struct {
	repository user.Repository
}

func NewAuthService(repository user.Repository) *AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) Authenticate(req *LoginRequest, secretKey string) (string, error) {
	accessToken, err := user.ValidateUser(s.repository, req.Email, req.Password, secretKey)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (s *AuthService) Signup(req *SignupRequest) error {
	return user.Signup(s.repository, req.Email, req.Password, req.FirstName, req.LastName)
}
