package entity

import "time"

type Task struct {
	ID          string    `bson:"_id,omitempty"`
	Title       string    `bson:"title"`
	Status      string    `bson:"status"` // например, "completed"
	CompletedAt time.Time `bson:"completed_at"`
}
