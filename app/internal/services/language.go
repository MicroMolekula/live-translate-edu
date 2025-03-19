package services

import (
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/repository"
)

type LanguageService struct {
	languageRepository *repository.LanguageRepository
}

func NewLanguageService(languageRepository *repository.LanguageRepository) *LanguageService {
	return &LanguageService{languageRepository: languageRepository}
}

func (s *LanguageService) Create(title string, code string) error {
	return s.languageRepository.Create(title, code)
}

func (s *LanguageService) GetAll() ([]dto.Language, error) {
	languagesModels, err := s.languageRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return dto.CreateLanguagesArrayFromModels(languagesModels), err
}
