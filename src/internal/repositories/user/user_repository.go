package user_repository

import (
	"github.com/aliaydins/todo/internal/models/user"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Migration() error {
	err := r.db.AutoMigrate(&user.User{})
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Create(user *user.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByUserName(email string) (*user.User, error) {
	user := new(user.User)
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
