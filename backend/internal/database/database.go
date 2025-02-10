package database

import (
	"backend/internal/configs"
	"backend/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func Init() *gorm.DB {
	dbConn, err := gorm.Open(postgres.Open(makeConfigString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection error: %s", err)
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
		log.Fatalf("Database migration error: %s", err)
	}

	return dbConn
}

func GetDb() *gorm.DB {
	if db == nil {
		db = Init()
		sleep := time.Duration(1)
		for db == nil {
			sleep *= 2
			time.Sleep(sleep * time.Second)
			db = Init()
			if sleep >= 10 {
				break
			}
		}
	}
	return db
}

func makeConfigString() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		configs.Cfg.Database.Host,
		configs.Cfg.Database.Port,
		configs.Cfg.Database.User,
		configs.Cfg.Database.Password,
		configs.Cfg.Database.DbName,
		configs.Cfg.Database.Timezone,
	)
}
