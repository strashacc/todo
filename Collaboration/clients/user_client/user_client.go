package userclient

import (
	user_pb "collaboration/pb/user/generated"
	"context"
	"time"

	"google.golang.org/grpc"
)

func getClient() (*grpc.ClientConn,
	user_pb.UserServiceClient,
	context.Context,
	context.CancelFunc,
	error) {
	conn, err := grpc.Dial("team:50051", grpc.WithInsecure())
	if err != nil {
		return nil, nil, nil, nil, err
	}

	client := user_pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	return conn, client, ctx, cancel, nil
}

func GetUser(id uint64) (*user_pb.User, error) {
	conn, client, ctx, cancel, err := getClient()
	if err != nil {
		return &user_pb.User{}, err
	}
	defer conn.Close()
	defer cancel()
	res, err := client.GetUser(ctx, &user_pb.GetUserRequest{Id: id})
	if err != nil {
		return &user_pb.User{}, err
	}
	return res.User, nil
}