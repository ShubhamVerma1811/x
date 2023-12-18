package database

import (
	"github.com/ShubhamVerma1811/x/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func SetupDB() (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(sqlite.Open("x.db"), &gorm.Config{
		TranslateError: true,
		// Logger:         logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	if db.AutoMigrate(&model.Bookmark{}) != nil {
		return nil, err
	}

	return db, nil

}
