package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Role       string
	Surname    string
	Name       string
	Email      string `gorm:"unique"`
	Password   string
	GroupID    uint
	LanguageID uint
	Messages   []Message
	Lessons    []*Lesson `gorm:"many2many:user_lesson;"`
}

type Group struct {
	gorm.Model
	Title string
	Code  string `gorm:"unique"`
	Users []User
}

type Language struct {
	gorm.Model
	Code            string
	Title           string
	Users           []User
	MessageContents []MessageContent
	LessonContents  []LessonContent
}

type Message struct {
	gorm.Model
	UserID          uint
	MessageContents []MessageContent
	LessonID        uint
}

type MessageContent struct {
	gorm.Model
	MessageID  uint
	Content    string
	LanguageID uint
}

type Lesson struct {
	gorm.Model
	Presentation   string
	DateTimeStart  time.Time
	NumberRoom     uint
	Messages       []Message
	Users          []*User `gorm:"many2many:user_lesson;"`
	LessonContents []LessonContent
}

type LessonContent struct {
	gorm.Model
	Theme      string
	Content    string
	LessonID   uint
	LanguageID uint
}
