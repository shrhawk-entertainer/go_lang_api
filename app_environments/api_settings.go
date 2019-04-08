package app_environments

import (
	"github.com/jinzhu/gorm"
	"github.com/vsouza/go-gin-boilerplate/models"
)

func MigrateDb(db *gorm.DB){
	db.AutoMigrate(models.GormUser{})
}
