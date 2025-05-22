package grpc

import (
	pb "collaboration/pb/team_pb"
)

type Server struct {
	pb.UnimplementedTeamServiceServer
}

func NewServer() *Server {
	return &Server{}
}