package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"timekeeper-backend/internal/remote"
)

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "timekeeper.db")
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&remote.Remote{})

	return db
}
