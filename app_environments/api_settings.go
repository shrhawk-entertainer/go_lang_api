package app_environments

import (
	"github.com/jinzhu/gorm"
	"github.com/shrhawk-entertainer/go_lang_api/models"
)

func MigrateDb(db *gorm.DB){
	db.AutoMigrate(models.GormUser{})
}
