package services

import (
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/repository"
)

type LessonService struct {
	lessonRepository *repository.LessonRepository
}

func NewLessonService(lessonRepository *repository.LessonRepository) *LessonService {
	return &LessonService{
		lessonRepository: lessonRepository,
	}
}

func (ls *LessonService) CreateLesson(lesson *dto.LessonCreate) error {
	return ls.lessonRepository.Create(lesson)
}
