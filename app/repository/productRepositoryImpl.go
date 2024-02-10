package repository

import (
	"api-exemple/app/model"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const createProduct = `INSERT INTO products (description, price, active, type) VALUES ($1, $2, $3, $4) RETURNING id`
const updateProduct = `UPDATE products set description = $2, price = $3, active = $4, type = $5 WHERE id = $1`
const deleteProduct = `DELETE FROM products WHERE id = $1`

const getProduct = `SELECT id, description, price, active, type FROM products WHERE id = $1 LIMIT 1`
const listProducts = `SELECT id, description, price, active, type FROM products ORDER BY description`

type ProductRepositoryImpl struct {
	Db *pgx.Conn
}

// Delete implements ProductRepository.
func (p *ProductRepositoryImpl) Delete(id uuid.UUID) error {
	_, err := p.Db.Exec(context.Background(), deleteProduct, id)
	return err
}

// FindAll implements ProductRepository.
func (p *ProductRepositoryImpl) FindAll() ([]model.Product, error) {
	rows, err := p.Db.Query(context.Background(), listProducts)
	if err != nil {
		return nil, err
	}
	var items = make([]model.Product, 0)
	for rows.Next() {
		var i model.Product
		if err := rows.Scan(
			&i.ID,
			&i.Description,
			&i.Price,
			&i.Active,
			&i.Type,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// FindById implements ProductRepository.
func (p *ProductRepositoryImpl) FindById(id uuid.UUID) (model.Product, error) {
	row := p.Db.QueryRow(context.Background(), getProduct, id)
	var product model.Product
	err := row.Scan(
		&product.ID,
		&product.Description,
		&product.Price,
		&product.Active,
		&product.Type,
	)
	return product, err
}

// Save implements ProductRepository.
func (p *ProductRepositoryImpl) Save(product model.Product) (uuid.UUID, error) {
	row := p.Db.QueryRow(context.Background(), createProduct,
		product.Description,
		product.Price,
		product.Active,
		product.Type,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

// Update implements ProductRepository.
func (p *ProductRepositoryImpl) Update(product model.Product) error {
	_, err := p.Db.Exec(context.Background(), updateProduct,
		product.ID,
		product.Description,
		product.Price,
		product.Active,
		product.Type,
	)
	return err
}

func NewProductRepositoryImpl(Db *pgx.Conn) ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}
