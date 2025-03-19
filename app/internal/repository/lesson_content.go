package repository

import (
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/models"
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

func (lcr *LessonContentRepository) GetLessonContentByLessonId(id uint) ([]*dto.LessonContent, error) {
	var lessonsContentModels []*models.LessonContent
	err := lcr.db.Model(&models.LessonContent{}).Where("lesson_id = ?", id).Find(&lessonsContentModels).Error
	if err != nil {
		return nil, err
	}
	return dto.ArrayLessonContentByModels(lessonsContentModels), nil
}

func (lcr *LessonContentRepository) GetLessonContentByLessonIds(ids []uint) ([]*dto.LessonContent, error) {
	var lessonsContentModels []*models.LessonContent
	err := lcr.db.Model(&models.LessonContent{}).Where("lesson_id in (?)", ids).Find(&lessonsContentModels).Error
	if err != nil {
		return nil, err
	}
	return dto.ArrayLessonContentByModels(lessonsContentModels), nil
}
