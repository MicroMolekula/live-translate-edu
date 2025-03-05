package repository

import (
	"github.com/live-translate-edu/internal/dto"
	"gorm.io/gorm"
)

type LessonContentRepository struct {
	db *gorm.DB
}

func NewLessonContentRepository(db *gorm.DB) *LessonContentRepository {
	return &LessonContentRepository{
		db: db,
	}
}

func (lcr *LessonContentRepository) Create(lessonContent *dto.LessonContent) error {
	lessonModel := dto.LessonContentToModels(lessonContent)
	return lcr.db.Create(lessonModel).Error
}
