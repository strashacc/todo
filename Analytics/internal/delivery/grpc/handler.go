package grpc

import (
	"context"

	pb "analytics/internal/delivery/grpc/pb"
	"analytics/internal/usecase"

	"google.golang.org/protobuf/types/known/emptypb"
)

type AnalyticsHandler struct {
	pb.UnimplementedAnalyticsServiceServer
	uc usecase.TaskUsecase
}

func NewAnalyticsHandler(uc usecase.TaskUsecase) *AnalyticsHandler {
	return &AnalyticsHandler{uc: uc}
}

func (h *AnalyticsHandler) GetTaskStats(ctx context.Context, _ *emptypb.Empty) (*pb.TaskStatsResponse, error) {
	stats, err := h.uc.GetStats()
	if err != nil {
		return nil, err
	}
	return convertToProto(stats), nil
}

func (h *AnalyticsHandler) GetTeamTaskStats(ctx context.Context, req *pb.TeamTaskStatsRequest) (*pb.TaskStatsResponse, error) {
	stats, err := h.uc.GetTeamStats(req.TeamId)
	if err != nil {
		return nil, err
	}
	return convertToProto(stats), nil
}

func (h *AnalyticsHandler) GetUserTaskStats(ctx context.Context, req *pb.UserTaskStatsRequest) (*pb.TaskStatsResponse, error) {
	stats, err := h.uc.GetUserStats(req.UserId)
	if err != nil {
		return nil, err
	}
	return convertToProto(stats), nil
}

func convertToProto(stats *usecase.Stats) *pb.TaskStatsResponse {
	return &pb.TaskStatsResponse{
		TotalTasks:      stats.TotalTasks,
		CompletedTasks:  stats.CompletedTasks,
		InProgressTasks: stats.InProgress,
		OverdueTasks:    stats.Overdue,
	}
}
