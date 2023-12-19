package database

import (
	"fmt"

	"github.com/ShubhamVerma1811/x/model"

	"os"
	"path/filepath"

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

	xDir := filepath.Join(os.Getenv("HOME"), ".config", "x")
	dbPath := filepath.Join(xDir, "x.db")

	if _, err := os.Stat(xDir); os.IsNotExist(err) {
		fmt.Println("Creating the database at", dbPath)
		os.MkdirAll(xDir, os.ModePerm)
	}

	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
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
