package repository

import (
	"github.com/live-translate-edu/internal/models"
	"gorm.io/gorm"
)

type LanguageRepository struct {
	db *gorm.DB
}

func NewLanguageRepository(db *gorm.DB) *LanguageRepository {
	return &LanguageRepository{
		db: db,
	}
}

func (lr *LanguageRepository) Create(title string, code string) error {
	return lr.db.Create(&models.Language{
		Title: title,
		Code:  code,
	}).Error
}
