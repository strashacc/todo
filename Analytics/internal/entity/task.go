package entity

import "time"

type Task struct {
	ID          string    `bson:"_id,omitempty"`
	Title       string    `bson:"title"`
	Status      string    `bson:"status"`       // например, "completed", "pending" и т.п.
	CompletedAt time.Time `bson:"completed_at"` // можно оставить для хранения даты выполнения
	Deadline    time.Time `bson:"deadline"`     // ОБЯЗАТЕЛЬНО добавить поле для дедлайна (чтобы считать просрочено!)
}
