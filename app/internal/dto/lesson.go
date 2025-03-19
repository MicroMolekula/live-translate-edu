package dto

import (
	"fmt"
	"github.com/live-translate-edu/internal/utils"
	"strings"
	"time"
)

type Lesson struct {
	Presentation  string   `json:"-"`
	DateTimeStart DateTime `json:"date_start"`
	NumberRoom    string   `json:"number_room"`
}

type LessonFull struct {
	Id       int              `json:"id"`
	Code     string           `json:"code"`
	Contents []*LessonContent `json:"contents"`
	Teacher  string           `json:"teacher"`
	Lesson
}

type LessonCreate struct {
	Theme          string   `json:"theme"`
	ThemeTranslate string   `json:"-"`
	GroupId        int      `json:"group_id"`
	LanguageCodes  []string `json:"languages_codes"`
	DateTimeStart  DateTime `json:"date_start"`
	NumberRoom     string   `json:"number_room"`
	TeacherId      int      `json:"teacher_id"`
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

func (t *DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(*t).Format("02-01-2006 15:04:05"))), nil
}

func (t *DateTime) ToTime() time.Time {
	return time.Time(*t)
}
