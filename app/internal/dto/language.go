package dto

import "github.com/live-translate-edu/internal/models"

type Language struct {
	Code  string `json:"code" binding:"required"`
	Title string `json:"title" binding:"required"`
}

func CreateLanguageFromModel(model *models.Language) *Language {
	return &Language{
		Code:  model.Code,
		Title: model.Title,
	}
}

func CreateLanguagesArrayFromModels(models []*models.Language) []Language {
	result := make([]Language, len(models))
	for i, model := range models {
		result[i] = *CreateLanguageFromModel(model)
	}
	return result
}
