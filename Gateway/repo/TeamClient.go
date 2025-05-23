package repo

import (
	"context"
	team_pb "gateway/pb/team_pb"
	"time"

	"google.golang.org/grpc"
)

func getClient() (*grpc.ClientConn,
				  team_pb.TeamServiceClient, 
				  context.Context, 
				  context.CancelFunc, 
				  error) {
	conn, err := grpc.Dial("team:50051", grpc.WithInsecure())
	if err != nil {
		return nil, nil, nil, nil, err
	}

	client := team_pb.NewTeamServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	return conn, client, ctx, cancel, nil
}

func DeleteTeam(id string) (*team_pb.TeamResponse, error) {
	conn, client, ctx, cancel, err := getClient()
	if err != nil {
		return &team_pb.TeamResponse{Success: false, Message: "Internal server error"}, err
	}
	defer conn.Close()
	defer cancel()
	res, err := client.DeleteTeam(ctx, &team_pb.DeleteTeamRequest{TeamId: id})
	if err != nil {
		return &team_pb.TeamResponse{Success: false, Message: "Couldn't delete team"}, nil
	}
	return res, err
}