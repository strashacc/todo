package grpc

import (
	pb "analytics/internal/delivery/grpc/pb"
	"analytics/internal/usecase"
	"context"
)

type TaskHandler struct {
	pb.UnimplementedTaskServiceServer
	uc usecase.TaskUsecase
}

func NewTaskHandler(uc usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{uc: uc}
}

func (h *TaskHandler) GetStats(ctx context.Context, _ *pb.Empty) (*pb.TaskStats, error) {
	stats, err := h.uc.GetStats()
	if err != nil {
		return nil, err
	}
	return &pb.TaskStats{
		CompletedToday: stats.CompletedToday,
		CompletedWeek:  stats.CompletedWeek,
		TotalTasks:     stats.TotalTasks,
	}, nil
}
