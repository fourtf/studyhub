package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"type:varchar(255);unique_index"`
	Password string `json:"Password"`
}
