package repository

import "github.com/live-translate-edu/internal/models"

type IRepository interface {
	FindOneById(id uint) (models.IModel, error)
}
