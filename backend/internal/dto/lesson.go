package dto

import (
	"github.com/live-translate-edu/internal/utils"
	"strings"
	"time"
)

type Lesson struct {
	Presentation  string    `json:"-"`
	DateTimeStart time.Time `json:"date_start"`
	NumberRoom    int       `json:"number_room"`
}

type LessonCreate struct {
	Theme         string   `json:"theme"`
	GroupId       int      `json:"group_id"`
	LanguageCodes []string `json:"languages_codes"`
	DateTimeStart DateTime `json:"date_start"`
	NumberRoom    int      `json:"number_room"`
	TeacherId     int      `json:"teacher_id"`
}

type DateTime time.Time

func (t *DateTime) UnmarshalJSON(b []byte) error {
	timeString := strings.Trim(string(b), "\"")
	dateTime, err := utils.ParseDate(timeString)
	if err != nil {
		return err
	}
	*t = DateTime(dateTime)
	return nil
}

func (t *DateTime) ToTime() time.Time {
	return time.Time(*t)
}
