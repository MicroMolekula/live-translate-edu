package repository

import (
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/models"
	"github.com/live-translate-edu/internal/utils"
	"gorm.io/gorm"
)

type LessonRepository struct {
	db *gorm.DB
}

func NewLessonRepository(db *gorm.DB) *LessonRepository {
	return &LessonRepository{db: db}
}

func (lr *LessonRepository) Create(lesson *dto.LessonCreate) error {
	return lr.db.Transaction(func(tx *gorm.DB) error {
		var lessonModel models.Lesson
		err := tx.Create(&models.Lesson{
			GroupID:       uint(lesson.GroupId),
			TeacherID:     uint(lesson.TeacherId),
			NumberRoom:    lesson.NumberRoom,
			DateTimeStart: lesson.DateTimeStart.ToTime(),
			CodeRoom:      utils.GenerateCodeRoom(lesson.NumberRoom),
		}).Scan(&lessonModel).Error
		if err != nil {
			return err
		}
		var languages []models.Language
		err = tx.Model(models.Language{}).
			Where("code in (?)", lesson.LanguageCodes).
			Find(&languages).Error
		if err != nil {
			return err
		}
		lessonContents := make([]models.LessonContent, len(languages))
		for i, language := range languages {
			theme := lesson.Theme
			if language.Code != "ru" {
				theme = lesson.ThemeTranslate
			}
			lessonContents[i] = models.LessonContent{
				Theme:      theme,
				LanguageID: language.ID,
				LessonID:   lessonModel.ID,
				Content:    "",
			}
		}
		if err := tx.Create(&lessonContents).Error; err != nil {
			return err
		}
		return nil
	})
}

func (lr *LessonRepository) GetAll() ([]*models.Lesson, error) {
	var lessons []*models.Lesson
	if err := lr.db.Find(&lessons).Error; err != nil {
		return nil, err
	}
	return lessons, nil
}

func (lr *LessonRepository) GetByUserId(id uint) ([]*models.Lesson, error) {
	var lessons []*models.Lesson
	if err := lr.db.Model(&models.Lesson{}).Where("teacher_id=?", id).Find(&lessons).Error; err != nil {
		return nil, err
	}
	return lessons, nil
}

func (lr *LessonRepository) GetByGroupId(groupId int) ([]*models.Lesson, error) {
	var lessons []*models.Lesson
	if err := lr.db.Model(&models.Lesson{}).Where("group_id=?", groupId).Find(&lessons).Error; err != nil {
		return nil, err
	}
	return lessons, nil
}
