package models

// Voter represents each registered voter
type Voter struct {
	ID      uint   `gorm:"primaryKey"`
	NIM     string `gorm:"unique;not null"`
	QRImage string `gorm:"type:text"` // store QR code (base64)
}
