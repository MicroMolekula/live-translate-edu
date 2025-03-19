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

func (lr *LanguageRepository) GetAll() ([]*models.Language, error) {
	var languages []*models.Language
	if err := lr.db.Find(&languages).Error; err != nil {
		return nil, err
	}
	return languages, nil
}

func (lr *LanguageRepository) GetByIds(ids []uint) (map[uint]*models.Language, error) {
	var languages []*models.Language
	if err := lr.db.Model(&models.Language{}).Where("id in (?)", ids).Find(&languages).Error; err != nil {
		return nil, err
	}
	result := make(map[uint]*models.Language)
	for _, language := range languages {
		result[language.ID] = language
	}
	return result, nil
}
