package repository

import (
	"api-exemple/app/model"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const createUser = `INSERT INTO users (name, email, password, active) VALUES ($1, $2, $3, $4) RETURNING id`
const updateUser = `UPDATE users set name = $2, email = $3, password = $4, active = $5 WHERE id = $1`
const deleteUser = `DELETE FROM users WHERE id = $1`

const getUser = `SELECT id, name, email, password, active FROM users WHERE id = $1 LIMIT 1`
const listUsers = `SELECT id, name, email, password, active FROM users ORDER BY name`
const getUserEmailAndPassword = `SELECT id, name, email, password, active FROM users WHERE email = $1 AND password = $2 LIMIT 1`
const getEmailRegistered = `SELECT id, name, email, password, active FROM users WHERE email = $1 LIMIT 1`

type UserRepositoryImpl struct {
	Db *pgx.Conn
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(id uuid.UUID) error {
	_, err := u.Db.Exec(context.Background(), deleteUser, id)
	return err
}

// FindAll implements UserRepository.
func (u *UserRepositoryImpl) FindAll() ([]model.User, error) {
	rows, err := u.Db.Query(context.Background(), listUsers)
	if err != nil {
		return nil, err
	}
	var items = make([]model.User, 0)
	for rows.Next() {
		var i model.User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.Active,
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

// FindByEmailAndPassword implements UserRepository.
func (u *UserRepositoryImpl) FindByEmailAndPassword(email string, password string) (model.User, error) {
	row := u.Db.QueryRow(context.Background(), getUserEmailAndPassword, email, password)
	var user model.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Active,
	)
	return user, err
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(id uuid.UUID) (model.User, error) {
	row := u.Db.QueryRow(context.Background(), getUser, id)
	var user model.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Active,
	)
	return user, err
}

// FindEmailRegistered implements UserRepository.
func (u *UserRepositoryImpl) FindEmailRegistered(email string) (model.User, error) {
	row := u.Db.QueryRow(context.Background(), getEmailRegistered, email)
	var user model.User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Active,
	)
	return user, err
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(user model.User) (uuid.UUID, error) {
	row := u.Db.QueryRow(context.Background(), createUser,
		user.Name,
		user.Email,
		user.Password,
		user.Active,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(user model.User) error {
	_, err := u.Db.Exec(context.Background(), updateUser,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.Active,
	)
	return err
}

func NewUserRepositoryImpl(Db *pgx.Conn) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}
