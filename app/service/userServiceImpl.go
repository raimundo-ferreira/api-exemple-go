package service

import (
	"api-exemple/app/data/request"
	"api-exemple/app/model"
	"api-exemple/app/repository"
	"api-exemple/pkg/utils"

	"github.com/google/uuid"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserSerive {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

// FindByEmailAndPassword implements UserSerive.
func (u *UserServiceImpl) FindByEmailAndPassword(login request.Login) (model.User, error) {
	user, err := u.UserRepository.FindByEmailAndPassword(login.Username, utils.Encrypt(login.Password))
	return user, err
}

// FindById implements UserSerive.
func (u *UserServiceImpl) FindById(id string) (model.User, error) {
	idUser, err := uuid.Parse(id)
	if err != nil {
		return model.User{}, err
	}

	user, err := u.UserRepository.FindById(idUser)
	return user, err
}

// Save implements UserSerive.
func (u *UserServiceImpl) Create(arg request.CreateUser) (model.User, error) {
	user, err := utils.TypeConverter[model.User](&arg)
	if err != nil {
		return *user, err
	}

	user.Password = utils.Encrypt(user.Password)
	id, err := u.UserRepository.Save(*user)
	if err == nil {
		user.ID = id
	}
	return *user, err
}

// Update implements UserSerive.
func (u *UserServiceImpl) Update(user model.User) error {
	_, err := u.UserRepository.FindById(user.ID)
	if err != nil {
		return err
	}

	user.Password = utils.Encrypt(user.Password)
	err = u.UserRepository.Update(user)
	return err
}

// Delete implements UserSerive.
func (u *UserServiceImpl) Delete(id string) error {
	idUser, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = u.UserRepository.FindById(idUser)
	if err != nil {
		return err
	}

	err = u.UserRepository.Delete(idUser)
	return err
}

// FindAll implements UserSerive.
func (u *UserServiceImpl) FindAll() ([]model.User, error) {
	users, err := u.UserRepository.FindAll()
	return users, err
}

// EmailRegistered implements UserSerive.
func (u *UserServiceImpl) FindEmailRegistered(email string) (model.User, error) {
	user, err := u.UserRepository.FindEmailRegistered(email)
	return user, err
}
