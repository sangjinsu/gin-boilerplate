package repository

import (
	"gin-boiler-plate/model"
	"gin-boiler-plate/repository"
)

type UserRepository interface {
	repository.Repository[model.User]
}

type userRepository struct {
	repository.Repository[model.User]
}

func NewUserRepository() UserRepository {
	return &userRepository{Repository: repository.NewRepository[model.User]()}
}
