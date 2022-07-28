package migrations

import (
	"github.com/Gabrielfr-bld/books-api-golang/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB)  {
	db.AutoMigrate(models.Book{})
}