package repository

import (
	"github.com/live-translate-edu/internal/database"
	"github.com/live-translate-edu/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: database.GetDb()}
}

func (ur *UserRepository) FindAll() ([]*models.User, error) {
	var users []*models.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) Create(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) FindOneById(id uint) (models.IModel, error) {
	var user *models.User
	if err := ur.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) FindOneByEmail(email string) (*models.User, error) {
	var user *models.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
