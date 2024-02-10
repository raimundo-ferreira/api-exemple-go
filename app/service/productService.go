package service

import (
	"api-exemple/app/data/request"
	"api-exemple/app/model"
)

type ProductSerive interface {
	Create(arg request.CreateProduct) (model.Product, error)
	Update(product model.Product) error
	Delete(id string) error
	FindById(id string) (model.Product, error)
	FindAll() ([]model.Product, error)
}
