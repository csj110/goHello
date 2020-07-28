package models

import "github.com/jinzhu/gorm"

type (
	LoginDto struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}
	User struct {
		gorm.Model
		Phone string `gorm:"unique_index;not null"`
	}
)