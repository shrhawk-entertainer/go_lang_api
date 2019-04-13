package models

import "github.com/jinzhu/gorm"

type GormUser struct {
	gorm.Model
	Name      string
	Email	  string `gorm:"unique;not null"`
	Gender    string
	Active    bool
}
