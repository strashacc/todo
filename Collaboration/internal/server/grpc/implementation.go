package grpc

import (
	"collaboration/internal/repo"
	pb "collaboration/pb/team_pb"
	"context"
	"errors"
	"log"
)

type Server struct {
	pb.UnimplementedTeamServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) CreateTeam(ctx context.Context, req *pb.CreateTeamRequest) (*pb.TeamResponse, error) {
	id, err := repo.CreateTeam(req)
	if errors.Is(err, &repo.BadUserError) {
		return &pb.TeamResponse{Success: false, Message: "Couldn't create team: Not every user exists"}, nil
	} else if err != nil {
		return &pb.TeamResponse{Success: false, Message: "Couldn't create team"}, err
	}
	return &pb.TeamResponse{Success: true, Id: id, Message: "Team successfully created"}, nil
}

func (s *Server) GetTeam(ctx context.Context, req *pb.GetTeamRequest) (*pb.Team, error) {
	team, err := repo.GetTeam(req)
	if err == &repo.NoDocumentError {
		log.Println("Database returned no documents")
		return &pb.Team{}, nil
	} else if err != nil {
		log.Println(err.Error())
		return &pb.Team{}, err
	}
	return team, nil
}

func (s *Server) GetTeams(ctx context.Context, req *pb.Empty) (*pb.TeamList, error) {
	return &pb.TeamList{}, nil
}

func (s *Server) UpdateTeam(ctx context.Context, req *pb.UpdateTeamRequest) (*pb.TeamResponse, error) {
	return &pb.TeamResponse{}, nil
}

func (s *Server) DeleteTeam(ctx context.Context, req *pb.DeleteTeamRequest) (*pb.TeamResponse, error) {
	res, err := repo.DeleteTeam(req)
	if errors.Is(err, &repo.NoDocumentError) {
		return &pb.TeamResponse{Success: false, Message: "Document doesn't exist"}, nil
	} else if errors.Is(err, &repo.DBWriteError) {
		return &pb.TeamResponse{Success: false, Message: "Deletion status unknown"}, nil
	} else if err != nil {
		log.Println(err.Error())
		return res, err
	}
	return res, nil
}