package service

import (
	"api-exemple/app/data/request"
	"api-exemple/app/model"
)

type UserSerive interface {
	Create(arg request.CreateUser) (model.User, error)
	Update(user model.User) error
	Delete(id string) error
	FindById(id string) (model.User, error)
	FindAll() ([]model.User, error)
	FindByEmailAndPassword(login request.Login) (model.User, error)
	FindEmailRegistered(email string) (model.User, error)
}
