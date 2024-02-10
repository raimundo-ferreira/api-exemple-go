package repository

import (
	"api-exemple/app/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(user model.User) (uuid.UUID, error)
	Update(user model.User) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (model.User, error)
	FindAll() ([]model.User, error)
	FindByEmailAndPassword(email, password string) (model.User, error)
	FindEmailRegistered(email string) (model.User, error)
}
