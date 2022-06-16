package user

import (
	"time"

	"github.com/aliaydins/todo/pkg/hash"
	jwt_helper "github.com/aliaydins/todo/pkg/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func ValidateUser(r Repository, email string, password string, secretKey string) (string, error) {
	user, err := r.FindByUserName(email)
	if err != nil {
		return "", ErrUserNotFound
	}
	if user.Password != hash.GenerateHash(password, user.Salt) {
		return "", ErrPasswordIncorrect
	}

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"email":     user.Email,
		"firstname": user.FirstName,
		"lastname":  user.LastName,
		"iat":       time.Now().Unix(),
		"exp": time.Now().Add(1 *
			time.Hour).Unix(),
	})

	accessToken := jwt_helper.GenerateToken(jwtClaims, secretKey)
	return accessToken, nil
}

func Signup(r Repository, email, password, firstname, lastname string) error {

	u, _ := r.FindByUserName(email)
	if u != nil {
		return ErrUserExist
	}

	salt := hash.GenerateSalt()
	user := &User{
		ID:        uuid.New(),
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		Salt:      salt,
		Password:  hash.GenerateHash(password, salt),
	}

	if err := r.Create(user); err != nil {
		return err
	}
	return nil
}

func Create(r Repository, u *User) error {
	return r.Create(u)
}
