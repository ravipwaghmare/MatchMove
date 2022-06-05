package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Connect to local database
func Connect() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
}
