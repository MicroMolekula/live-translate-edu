package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/live-translate-edu/internal/configs"
	"github.com/live-translate-edu/internal/dto"
	"github.com/live-translate-edu/internal/models"
	"github.com/live-translate-edu/internal/repository"
	"github.com/live-translate-edu/internal/services/speech_translate"
	"github.com/live-translate-edu/internal/utils/roles"
	"golang.org/x/exp/slices"
	"log"
)

type LessonService struct {
	lessonRepository        *repository.LessonRepository
	lessonContentRepository *repository.LessonContentRepository
	translateService        *speech_translate.TranslateServ
	userRepository          *repository.UserRepository
	languageRepository      *repository.LanguageRepository
}

func NewLessonService(
	lessonRepository *repository.LessonRepository,
	lessonContentRepository *repository.LessonContentRepository,
	cfg *configs.Config,
	userRepository *repository.UserRepository,
	languageRepository *repository.LanguageRepository) *LessonService {
	translateService, err := speech_translate.NewTranslateServ(cfg)
	if err != nil {
		log.Println(err)
	}
	return &LessonService{
		lessonRepository:        lessonRepository,
		translateService:        translateService,
		lessonContentRepository: lessonContentRepository,
		userRepository:          userRepository,
		languageRepository:      languageRepository,
	}
}

func (ls *LessonService) CreateLesson(lesson *dto.LessonCreate) error {
	if ls.translateService == nil {
		return errors.New("translateService is nil")
	}
	translateTheme, err := ls.translateService.TranslateText(context.Background(), lesson.Theme, &dto.TranslateLanguagesDto{
		Source: "ru",
		Target: "en",
	})
	if err != nil {
		return err
	}
	lesson.ThemeTranslate = translateTheme
	return ls.lessonRepository.Create(lesson)
}

func (ls *LessonService) GetAll() ([]*dto.LessonFull, error) {
	lessonsModels, err := ls.lessonRepository.GetAll()
	if err != nil {
		return nil, err
	}
	result, err := ls.createFullLesson(lessonsModels)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ls *LessonService) GetByUser(user *dto.UserDTO) ([]*dto.LessonFull, error) {
	var lessonsModels []*models.Lesson
	var err error
	switch user.Role {
	case roles.Teacher:
		lessonsModels, err = ls.lessonRepository.GetByUserId(user.Id)
		if err != nil {
			return nil, err
		}
	case roles.Student:
		lessonsModels, err = ls.lessonRepository.GetByGroupId(user.GroupId)
		if err != nil {
			return nil, err
		}
	}
	result, err := ls.createFullLesson(lessonsModels)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ls *LessonService) createFullLesson(lessonsModels []*models.Lesson) ([]*dto.LessonFull, error) {
	lessonsIds := make([]uint, len(lessonsModels))
	for i, lessonModel := range lessonsModels {
		lessonsIds[i] = lessonModel.ID
	}
	teacherIds := make([]uint, len(lessonsModels))
	for i, lessonsModel := range lessonsModels {
		teacherIds[i] = lessonsModel.TeacherID
	}
	teachers, err := ls.userRepository.FindByIds(teacherIds)
	if err != nil {
		return nil, err
	}
	lessonContents, err := ls.lessonContentRepository.GetLessonContentByLessonIds(lessonsIds)
	if err != nil {
		return nil, err
	}
	lessons := make(map[int]*dto.LessonFull, len(lessonsIds))
	for i, lessonModel := range lessonsModels {
		lessons[int(lessonsModels[i].ID)] = &dto.LessonFull{
			Id:       int(lessonsModels[i].ID),
			Code:     lessonModel.CodeRoom,
			Contents: make([]*dto.LessonContent, 0),
			Teacher:  fmt.Sprintf("%s %s", teachers[lessonModel.TeacherID].Surname, teachers[lessonModel.TeacherID].Name),
			Lesson: dto.Lesson{
				DateTimeStart: dto.DateTime(lessonModel.DateTimeStart),
				NumberRoom:    lessonModel.NumberRoom,
			},
		}
	}
	languagesIds := make([]uint, 0, len(lessonContents))
	for _, lessonContent := range lessonContents {
		if !slices.Contains(languagesIds, lessonContent.LanguageId) {
			languagesIds = append(languagesIds, lessonContent.LanguageId)
		}
	}
	languages, err := ls.languageRepository.GetByIds(languagesIds)
	if err != nil {
		return nil, err
	}
	for _, lessonContent := range lessonContents {
		lessonContent.LanguageCode = languages[lessonContent.LanguageId].Code
		lessons[int(lessonContent.LessonId)].Contents = append(lessons[int(lessonContent.LessonId)].Contents, lessonContent)
	}
	result := make([]*dto.LessonFull, 0, len(lessonsIds))
	for _, lesson := range lessons {
		result = append(result, lesson)
	}
	return result, nil
}
