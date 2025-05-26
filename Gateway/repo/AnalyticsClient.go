package repo

import (
	"context"
	"gateway/pb"
	"time"

	"google.golang.org/grpc"
)

func getAnalyticsClient() (*grpc.ClientConn,
	pb.AnalyticsServiceClient,
	context.Context,
	context.CancelFunc,
	error) {
	conn, err := grpc.Dial("analytics:50051", grpc.WithInsecure())
	if err != nil {
		return nil, nil, nil, nil, err
	}
	client := pb.NewAnalyticsServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return conn, client, ctx, cancel, nil
}

func GetUserTaskStats(userID string) (*pb.TaskStatsResponse, error) {
	conn, client, ctx, cancel, err := getAnalyticsClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	defer cancel()

	res, err := client.GetUserTaskStats(ctx, &pb.UserTaskStatsRequest{UserId: userID})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetTeamTaskStats(teamID string) (*pb.TaskStatsResponse, error) {
	conn, client, ctx, cancel, err := getAnalyticsClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	defer cancel()

	res, err := client.GetTeamTaskStats(ctx, &pb.TeamTaskStatsRequest{TeamId: teamID})
	if err != nil {
		return nil, err
	}
	return res, nil
}
