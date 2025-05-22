package grpc

import (
	"collaboration/internal/repo"
	pb "collaboration/pb/team_pb"
	"context"
	"errors"
)

type Server struct {
	pb.UnimplementedTeamServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) CreateTeam(ctx context.Context, req *pb.CreateTeamRequest) (*pb.TeamResponse, error) {
	id, err := repo.CreateTeam(req)
	if errors.As(err, &repo.BadUserError{}) {
		return &pb.TeamResponse{Success: false, Message: "Couldn't create team: Not every user exists"}, nil
	}
	if err != nil {
		return &pb.TeamResponse{Success: false, Message: "Couldn't create team"}, err
	}
	return &pb.TeamResponse{Success: true, Id: id, Message: "Team successfully created"}, nil
}

func (s *Server) GetTeam(ctx context.Context, req *pb.GetTeamRequest) (*pb.Team, error) {
	return &pb.Team{}, nil
}

func (s *Server) GetTeams(ctx context.Context, req *pb.Empty) (*pb.TeamList, error) {
	return &pb.TeamList{}, nil
}

func (s *Server) UpdateTeam(ctx context.Context, req *pb.UpdateTeamRequest) (*pb.TeamResponse, error) {
	return &pb.TeamResponse{}, nil
}

func (s *Server) DeleteTeam(ctx context.Context, req *pb.DeleteTeamRequest) (*pb.TeamResponse, error) {
	res, err := repo.DeleteTeam(req)
	if errors.As(err, &repo.NoDocumentError{}) {
		return &pb.TeamResponse{Success: false, Message: "Document does not exist"}, nil
	} else if errors.As(err, &repo.DBWriteError{}) {
		return &pb.TeamResponse{Success: false, Message: "Deletion status unknown"}, nil
	} else if err != nil {
		return res, err
	}
	return res, err
}