package database

import (
	"github.com/jinzhu/gorm"
	"github.com/marcusziade/go-rest-api-practice-1/comment"
)

// Migrates the database and creates our comment table
func MigrateDatabase(database *gorm.DB) error {
	if result := database.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}

	return nil
}
