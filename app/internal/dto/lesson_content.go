package dto

import "github.com/live-translate-edu/internal/models"

type LessonContent struct {
	Theme      string
	Content    string
	LessonId   uint
	LanguageId uint
}

func LessonContentToModels(lessonContent *LessonContent) *models.LessonContent {
	return &models.LessonContent{
		Theme:      lessonContent.Theme,
		Content:    lessonContent.Content,
		LessonID:   lessonContent.LessonId,
		LanguageID: lessonContent.LanguageId,
	}
}
