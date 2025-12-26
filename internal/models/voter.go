package models

import "gorm.io/gorm"

type Voter struct {
	gorm.Model
	Name       string `json:"name" gorm:"type:varchar(100);not null"`
	Email      string `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Token      string `json:"token" gorm:"type:varchar(255);uniqueIndex"`      // Simpan SHA-256 (64 karakter)
	ShortToken string `json:"short_token" gorm:"type:varchar(10);uniqueIndex"` // Simpan Kode Pendek (6 karakter)
	HasVoted   bool   `json:"has_voted" gorm:"default:false"`
}
