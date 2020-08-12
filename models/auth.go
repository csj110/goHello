package models

import "github.com/jinzhu/gorm"

type (
	User struct {
		gorm.Model
		Phone string `gorm:"unique_index;not null"`
	}
)