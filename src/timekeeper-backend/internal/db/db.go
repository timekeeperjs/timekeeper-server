//go:test ignoretest

package db

import (
	"timekeeper-backend/internal/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "timekeeper.db")
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Remote{})

	return db
}
