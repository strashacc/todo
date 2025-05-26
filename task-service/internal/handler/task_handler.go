package handler

import (
	"context"
	"task-service/internal/db"
	"task-service/internal/model"
	"task-service/internal/repository"
	"task-service/internal/service"
	"task-service/pb/task_pb"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type taskHandler struct {
	task_pb.UnimplementedTaskServiceServer
	service service.TaskService
}

func NewTaskHandler() *taskHandler {
	repo := repository.NewTaskRepository(db.DB)
	svc := service.NewTaskService(repo)
	return &taskHandler{service: svc}
}

func (h *taskHandler) CreateTask(ctx context.Context, req *task_pb.CreateTaskRequest) (*task_pb.TaskResponse, error) {
	task := &model.Task{
		UserID:      req.UserId,
		TeamID:      req.TeamId,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		Deadline:    req.Deadline.AsTime(),
		Tags:        req.Tags,
		Folder:      req.Folder,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := h.service.Create(task)
	if err != nil {
		return &task_pb.TaskResponse{Success: false, Message: err.Error()}, nil
	}

	return &task_pb.TaskResponse{
		Success: true,
		Id:      uint64(task.ID),
		Message: "Task created successfully",
		Task:    convertToProto(task),
	}, nil
}

func (h *taskHandler) GetTask(ctx context.Context, req *task_pb.GetTaskRequest) (*task_pb.Task, error) {
	task, err := h.service.Get(req.TaskId)
	if err != nil {
		return nil, err
	}
	return convertToProto(task), nil
}

func (h *taskHandler) GetTasks(ctx context.Context, _ *task_pb.Empty) (*task_pb.TaskList, error) {
	tasks, err := h.service.List()
	if err != nil {
		return nil, err
	}
	var protoTasks []*task_pb.Task
	for _, task := range tasks {
		t := task
		protoTasks = append(protoTasks, convertToProto(&t))
	}
	return &task_pb.TaskList{Tasks: protoTasks}, nil
}

func (h *taskHandler) UpdateTask(ctx context.Context, req *task_pb.UpdateTaskRequest) (*task_pb.TaskResponse, error) {
	task := req.Update
	existing, err := h.service.Get(req.TaskId)
	if err != nil {
		return &task_pb.TaskResponse{Success: false, Message: "Task not found"}, nil
	}
	existing.Title = task.Title
	existing.Description = task.Description
	existing.Status = task.Status
	existing.Priority = task.Priority
	existing.Deadline = task.Deadline.AsTime()
	existing.Tags = task.Tags
	existing.Folder = task.Folder
	existing.UpdatedAt = time.Now()

	err = h.service.Update(existing)
	if err != nil {
		return &task_pb.TaskResponse{Success: false, Message: err.Error()}, nil
	}
	return &task_pb.TaskResponse{Success: true, Message: "Task updated", Task: convertToProto(existing)}, nil
}

func (h *taskHandler) DeleteTask(ctx context.Context, req *task_pb.DeleteTaskRequest) (*task_pb.TaskResponse, error) {
	err := h.service.Delete(req.TaskId)
	if err != nil {
		return &task_pb.TaskResponse{Success: false, Message: err.Error()}, nil
	}
	return &task_pb.TaskResponse{Success: true, Message: "Task deleted"}, nil
}

func convertToProto(task *model.Task) *task_pb.Task {
	return &task_pb.Task{
		Id:          uint64(task.ID),
		UserId:      task.UserID,
		TeamId:      task.TeamID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
		Deadline:    timestamppb.New(task.Deadline),
		Tags:        task.Tags,
		Folder:      task.Folder,
		CreatedAt:   timestamppb.New(task.CreatedAt),
		UpdatedAt:   timestamppb.New(task.UpdatedAt),
	}
}
