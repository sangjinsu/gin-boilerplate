package repository

import (
	"errors"
	"gin-boiler-plate/db"
	"gorm.io/gorm"
)

type Repository[model any] interface {
	FindById(uint) (*model, error)
	FindAll() ([]model, error)
	Create(model) (*model, error)
	Update(uint, model) (*model, error)
	Delete(uint) error
}

type repository[model any] struct {
	tx *gorm.DB
}

func NewRepository[model any]() *repository[model] {
	return &repository[model]{tx: db.GetDB()}
}

func (r *repository[model]) FindById(id uint) (*model, error) {
	newModel := new(model)
	if err := r.tx.First(&newModel, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return newModel, nil
}

func (r *repository[model]) FindAll() ([]model, error) {
	newModels := make([]model, 0)
	if err := r.tx.Find(&newModels).Error; err != nil {
		return nil, err
	}
	return newModels, nil
}

func (r *repository[model]) Create(createModel model) (*model, error) {
	if err := r.tx.Create(&createModel).Error; err != nil {
		return nil, err
	}
	return &createModel, nil
}

func (r *repository[model]) Update(id uint, updateModel model) (*model, error) {
	newModel := new(model)
	if err := r.tx.First(&newModel, id).Updates(&updateModel).Error; err != nil {
		return nil, err
	}
	return newModel, nil
}

func (r *repository[model]) Delete(id uint) error {
	newModel := new(model)
	if err := r.tx.Delete(&newModel, id).Error; err != nil {
		return err
	}
	return nil
}
