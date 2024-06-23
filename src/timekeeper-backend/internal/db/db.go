//go:test ignoretest

package db

import (
	"timekeeper-backend/internal/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Init(dbFilePath string) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbFilePath)
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Remote{})

	return db
}

func ClearDatabase(dbFilePath string) error {
	db, err := gorm.Open("sqlite3", dbFilePath)
	if err != nil {
		return err
	}
	defer db.Close()

	// Drop all tables
	err = db.DropTableIfExists(&models.Remote{}).Error
	if err != nil {
		return err
	}

	return nil
}
