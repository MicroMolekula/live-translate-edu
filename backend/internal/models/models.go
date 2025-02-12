package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type IModel interface {
	GetId() uint
}

type ArrayModels[T IModel] []*T

type User struct {
	gorm.Model
	Role       string
	Surname    string
	Name       string
	Email      string `gorm:"unique"`
	Password   string
	GroupID    sql.NullInt64
	LanguageID sql.NullInt64
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

func (u *User) GetId() uint            { return u.ID }
func (g *Group) GetId() uint           { return g.ID }
func (l *Language) GetId() uint        { return l.ID }
func (m *Message) GetId() uint         { return m.ID }
func (mc *MessageContent) GetId() uint { return mc.ID }
func (le *Lesson) GetId() uint         { return le.ID }
func (lc *LessonContent) GetId() uint  { return lc.ID }
