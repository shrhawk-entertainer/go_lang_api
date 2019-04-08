package db

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vsouza/go-gin-boilerplate/config"
)

var GlobalDb *gorm.DB

func Init() *gorm.DB {
	config := config.GetConfig()
	MySqlConnectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.GetString("mysql.user"),
		config.GetString("mysql.password"),
		config.GetString("mysql.host"),
		config.GetString("mysql.port"),
		config.GetString("mysql.database"),
	)
	db, DbError := gorm.Open("mysql", MySqlConnectionString)
	if DbError != nil {
		errors.New(DbError.Error())
	}
	GlobalDb = db
	return db
}

func GetDB() *gorm.DB {
	return GlobalDb
}


