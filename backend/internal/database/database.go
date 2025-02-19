package database

import (
	"errors"
	"fmt"
	"github.com/live-translate-edu/internal/configs"
	"github.com/live-translate-edu/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *configs.Config) (*gorm.DB, error) {
	dbConn, err := gorm.Open(postgres.Open(makeConfigString(cfg)), &gorm.Config{})
	if err != nil {
		return nil, errors.New(fmt.Sprint("database connection error: ", err))
	}

	err = dbConn.AutoMigrate(
		models.Group{},
		models.Language{},
		models.User{},
		models.Lesson{},
		models.LessonContent{},
		models.Message{},
		models.MessageContent{},
	)
	if err != nil {
		return nil, errors.New(fmt.Sprint("database migration error: ", err))
	}

	return dbConn, nil
}

func makeConfigString(cfg *configs.Config) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DbName,
		cfg.Database.Timezone,
	)
}
