package service

import (
	"gin-boiler-plate/model"
	repository "gin-boiler-plate/repository/user"
)

type UserService interface {
	FindById(uint) (*model.User, error)
	FindAll() ([]model.User, error)
	Create(model.User) (*model.User, error)
	Update(uint, model.User) (*model.User, error)
	Delete(uint) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService() UserService {
	return &userService{userRepository: repository.NewUserRepository()}
}

func (u *userService) FindById(id uint) (*model.User, error) {
	return u.userRepository.FindById(id)
}

func (u *userService) FindAll() ([]model.User, error) {
	return u.userRepository.FindAll()
}

func (u *userService) Create(createUser model.User) (*model.User, error) {
	return u.userRepository.Create(createUser)
}

func (u *userService) Update(id uint, updateUser model.User) (*model.User, error) {
	return u.userRepository.Update(id, updateUser)
}

func (u *userService) Delete(id uint) error {
	return u.userRepository.Delete(id)
}
