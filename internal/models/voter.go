package models

import "gorm.io/gorm"

type Voter struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(150);uniqueIndex;not null"`
	QRToken  string `gorm:"type:char(64);uniqueIndex;not null"` // SHA-256 hex
	HasVoted bool   `gorm:"default:false"`
}
