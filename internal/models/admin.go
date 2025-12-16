package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Password string `gorm:"type:char(60);not null"` // bcrypt hash
}
