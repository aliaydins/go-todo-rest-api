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

func (ur *userRepository) Migration() error {
	err := ur.db.Migrator().DropTable(&user.User{})
	if err != nil {
		return err
	}
	err = ur.db.AutoMigrate(&user.User{})
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) Create(user *user.User) error {
	return ur.db.Create(user).Error
}

func (ur *userRepository) FindByUserName(email string) (*user.User, error) {
	user := new(user.User)
	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
