package models

// Vote represents a voting record
type Vote struct {
	ID          uint   `gorm:"primaryKey"`
	VoterNIM    string `gorm:"not null"`
	CandidateID uint   `gorm:"not null"`
	Hash        string `gorm:"type:text;not null"`
}
