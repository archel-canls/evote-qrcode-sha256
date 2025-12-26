package models

import "time"

type Vote struct {
	ID          uint   `gorm:"primaryKey"`
	VoterID     uint   `gorm:"not null"`
	CandidateID uint   `gorm:"not null"`
	VoteHash    string `gorm:"uniqueIndex;not null"`
	CreatedAt   time.Time
}
