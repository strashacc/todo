package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	col *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) *TaskRepository {
	return &TaskRepository{
		col: db.Collection("tasks"),
	}
}

func (r *TaskRepository) CountStats(ctx context.Context) (completedToday, completedWeek, total int64, err error) {
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	weekAgo := now.AddDate(0, 0, -7)

	total, err = r.col.CountDocuments(ctx, bson.M{})
	if err != nil {
		return
	}

	completedToday, err = r.col.CountDocuments(ctx, bson.M{
		"status":       "completed",
		"completed_at": bson.M{"$gte": todayStart},
	})
	if err != nil {
		return
	}

	completedWeek, err = r.col.CountDocuments(ctx, bson.M{
		"status":       "completed",
		"completed_at": bson.M{"$gte": weekAgo},
	})
	return
}
