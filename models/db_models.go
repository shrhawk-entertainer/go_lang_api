package models

import "github.com/jinzhu/gorm"

type GormUser struct {
	gorm.Model
	Name      string
	Email	  string
	Gender    string
	Active    bool
}
