package models

import "gorm.io/gorm"

type Vote struct {
	gorm.Model

	VoterID     uint `gorm:"not null;index;uniqueIndex:idx_voter_once"`
	CandidateID uint `gorm:"not null;index"`

	VoteHash string `gorm:"type:char(64);uniqueIndex;not null"` // SHA-256 hex
}
