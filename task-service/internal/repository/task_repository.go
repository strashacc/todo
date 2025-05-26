package repository

import (
	"task-service/internal/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *model.Task) error
	GetByID(id uint64) (*model.Task, error)
	GetAll() ([]model.Task, error)
	Update(task *model.Task) error
	Delete(id uint64) error
}

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepo) GetByID(id uint64) (*model.Task, error) {
	var task model.Task
	err := r.db.First(&task, id).Error
	return &task, err
}

func (r *taskRepo) GetAll() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepo) Update(task *model.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepo) Delete(id uint64) error {
	return r.db.Delete(&model.Task{}, id).Error
}
