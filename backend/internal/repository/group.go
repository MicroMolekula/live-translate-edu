package repository

import (
	"github.com/live-translate-edu/internal/models"
	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (gr *GroupRepository) Create(group *models.Group) error {
	return gr.db.Create(group).Error
}

func (gr *GroupRepository) GetAll() ([]*models.Group, error) {
	var groups []*models.Group
	if err := gr.db.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (gr *GroupRepository) GetById(id int) (*models.Group, error) {
	var group *models.Group
	if err := gr.db.Preload("Users").Where("id = ?", id).First(&group).Error; err != nil {
		return nil, err
	}
	return group, nil
}
