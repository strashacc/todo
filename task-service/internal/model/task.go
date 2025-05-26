package model

import (
	"time"

	"github.com/lib/pq"
)

type Task struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint64
	TeamID      string
	Title       string
	Description string
	Status      string
	Priority    string
	Deadline    time.Time
	Tags        pq.StringArray `gorm:"type:text[]"`
	Folder      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
