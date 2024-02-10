package repository

import (
	"api-exemple/app/model"

	"github.com/google/uuid"
)

type ProductRepository interface {
	Save(product model.Product) (uuid.UUID, error)
	Update(product model.Product) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (model.Product, error)
	FindAll() ([]model.Product, error)
}
