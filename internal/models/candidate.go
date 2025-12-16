package models

import "time"

type Candidate struct {
	ID          uint      `gorm:"primaryKey" json:"ID"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	CreatedAt   time.Time `json:"CreatedAt"`
}
