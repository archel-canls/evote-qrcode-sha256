package models

// Candidate represents election candidates
type Candidate struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}
