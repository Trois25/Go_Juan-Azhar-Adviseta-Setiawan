package user

import "time"

type UserCore struct {
	ID        uint
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataInterface interface {
	Insert(data UserCore) error
	GetAllUsers()([]UserCore,error) 
	Login(email, password string) (UserCore, string, error)
}

type UseCaseInterface interface {
	GetAllUsers()([]UserCore,error) 
	Create(data UserCore) error
	Login(email, password string) (UserCore, string, error)
}
