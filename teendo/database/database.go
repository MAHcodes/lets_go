package database

import (
	"log"

	"github.com/MAHcodes/lets_go/teendo/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open(config.DB_URL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	DB = db
}

func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}
