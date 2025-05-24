package models

import (
	"time"

	"github.com/lib/pq"
)

type Task struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Status      string         `json:"status"`   // например: pending, in-progress, done
	Priority    string         `json:"priority"` // low, medium, high
	Deadline    time.Time      `json:"deadline"`
	Tags        pq.StringArray `gorm:"type:text[]" json:"tags"`
	Folder      string         `json:"folder"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
