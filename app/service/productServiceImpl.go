package service

import (
	"api-exemple/app/data/request"
	"api-exemple/app/model"
	"api-exemple/app/repository"
	"api-exemple/pkg/utils"

	"github.com/google/uuid"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
}

// Create implements ProductSerive.
func (p *ProductServiceImpl) Create(arg request.CreateProduct) (model.Product, error) {
	product, err := utils.TypeConverter[model.Product](&arg)
	if err != nil {
		return *product, err
	}

	id, err := p.ProductRepository.Save(*product)
	if err == nil {
		product.ID = id
	}
	return *product, err
}

// Delete implements ProductSerive.
func (p *ProductServiceImpl) Delete(id string) error {
	idProduct, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = p.ProductRepository.FindById(idProduct)
	if err != nil {
		return err
	}

	err = p.ProductRepository.Delete(idProduct)
	return err
}

// FindAll implements ProductSerive.
func (p *ProductServiceImpl) FindAll() ([]model.Product, error) {
	products, err := p.ProductRepository.FindAll()
	return products, err
}

// FindById implements ProductSerive.
func (p *ProductServiceImpl) FindById(id string) (model.Product, error) {
	idProduct, err := uuid.Parse(id)
	if err != nil {
		return model.Product{}, err
	}

	product, err := p.ProductRepository.FindById(idProduct)
	return product, err
}

// Update implements ProductSerive.
func (p *ProductServiceImpl) Update(product model.Product) error {
	_, err := p.ProductRepository.FindById(product.ID)
	if err != nil {
		return err
	}

	err = p.ProductRepository.Update(product)
	return err
}

func NewProductServiceImpl(ProductRepository repository.ProductRepository) ProductSerive {
	return &ProductServiceImpl{
		ProductRepository: ProductRepository,
	}
}
