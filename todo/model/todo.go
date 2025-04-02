package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
}
