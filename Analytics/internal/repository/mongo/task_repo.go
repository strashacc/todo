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

// Общая статистика по всем задачам
func (r *TaskRepository) CountStats(ctx context.Context) (total, completed, inProgress, overdue int64, err error) {
	return r.countByFilter(ctx, bson.M{})
}

// Статистика по задачам конкретной команды
func (r *TaskRepository) CountTeamStats(ctx context.Context, teamID string) (total, completed, inProgress, overdue int64, err error) {
	return r.countByFilter(ctx, bson.M{"team_id": teamID})
}

// Статистика по задачам конкретного пользователя
func (r *TaskRepository) CountUserStats(ctx context.Context, userID string) (total, completed, inProgress, overdue int64, err error) {
	return r.countByFilter(ctx, bson.M{"user_id": userID})
}

// Общая логика подсчета с учетом фильтра
func (r *TaskRepository) countByFilter(ctx context.Context, filter bson.M) (total, completed, inProgress, overdue int64, err error) {
	now := time.Now()

	// Всего задач
	total, err = r.col.CountDocuments(ctx, filter)
	if err != nil {
		return
	}

	// Выполнено
	filterCompleted := cloneFilter(filter)
	filterCompleted["status"] = "completed"
	completed, err = r.col.CountDocuments(ctx, filterCompleted)
	if err != nil {
		return
	}

	// Просрочено
	filterOverdue := cloneFilter(filter)
	filterOverdue["status"] = bson.M{"$ne": "completed"}
	filterOverdue["deadline"] = bson.M{"$lt": now}
	overdue, err = r.col.CountDocuments(ctx, filterOverdue)
	if err != nil {
		return
	}

	// В процессе
	filterInProgress := cloneFilter(filter)
	filterInProgress["status"] = bson.M{"$ne": "completed"}
	filterInProgress["deadline"] = bson.M{"$gte": now}
	inProgress, err = r.col.CountDocuments(ctx, filterInProgress)
	return
}

// Утилита клонирования фильтра
func cloneFilter(m bson.M) bson.M {
	clone := bson.M{}
	for k, v := range m {
		clone[k] = v
	}
	return clone
}
