package dto

import "github.com/live-translate-edu/internal/models"

type LessonContent struct {
	Theme        string `json:"theme"`
	Content      string `json:"content"`
	LessonId     uint   `json:"lesson_id"`
	LanguageId   uint   `json:"language_id"`
	LanguageCode string `json:"language_code"`
}

func LessonContentToModels(lessonContent *LessonContent) *models.LessonContent {
	return &models.LessonContent{
		Theme:      lessonContent.Theme,
		Content:    lessonContent.Content,
		LessonID:   lessonContent.LessonId,
		LanguageID: lessonContent.LanguageId,
	}
}

func LessonContentFromModels(lessonContent *models.LessonContent) *LessonContent {
	return &LessonContent{
		Theme:      lessonContent.Theme,
		Content:    lessonContent.Content,
		LessonId:   lessonContent.LessonID,
		LanguageId: lessonContent.LanguageID,
	}
}

func ArrayLessonContentByModels(lessonContents []*models.LessonContent) []*LessonContent {
	result := make([]*LessonContent, len(lessonContents))
	for i, lessonContent := range lessonContents {
		result[i] = LessonContentFromModels(lessonContent)
	}
	return result
}
