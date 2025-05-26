package service

import (
	"task-service/internal/model"
	"task-service/internal/repository"
)

type TaskService interface {
	Create(task *model.Task) error
	Get(id uint64) (*model.Task, error)
	List() ([]model.Task, error)
	Update(task *model.Task) error
	Delete(id uint64) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) Create(task *model.Task) error {
	return s.repo.Create(task)
}

func (s *taskService) Get(id uint64) (*model.Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) List() ([]model.Task, error) {
	return s.repo.GetAll()
}

func (s *taskService) Update(task *model.Task) error {
	return s.repo.Update(task)
}

func (s *taskService) Delete(id uint64) error {
	return s.repo.Delete(id)
}
