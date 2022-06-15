package user

type Repository interface {
	Migration() error
	Create(user *User) error
	FindByUserName(email string) (*User,error)
}