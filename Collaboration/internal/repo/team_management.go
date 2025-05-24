package repo

import (
	db "collaboration/db/mongodb"
	pb "collaboration/pb/team_pb"
	// user_pb "collaboration/pb/user/generated"
	"context"

	"github.com/google/uuid"
)

func CreateTeam(req *pb.CreateTeamRequest) (string, error) {
	col := db.GetDatabase().Collection("teams")

	if err := ValidateMembers(req.Members); err != nil {
		return "", err
	}

	id := uuid.New().String()
	insert := pb.Team{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
		Members:     req.Members,
	}

	res, err := col.InsertOne(context.TODO(), &insert)
	if err != nil {
		return "", err
	} else if !res.Acknowledged {
		return "", &unknownInsertError{}
	}
	return id, nil
}

func GetTeam(req *pb.GetTeamRequest) (*pb.Team, error) {
	col := db.GetDatabase().Collection("teams")

	var res *pb.Team
	if err := col.FindOne(context.TODO(), req).Decode(&res); err != nil {
		return &pb.Team{}, err
	}
	if res.Id == "" {
		return &pb.Team{}, &NoDocumentError
	}
	return res, nil
}

func GetTeams(req *pb.Empty) (*pb.TeamList, error) {
	// col := db.GetDatabase().Collection("teams")
	return &pb.TeamList{}, nil
}

func DeleteTeam(req *pb.DeleteTeamRequest) (*pb.TeamResponse, error) {
	col := db.GetDatabase().Collection("teams")

	if res, err := col.DeleteOne(context.TODO(), req); res.DeletedCount == 0 {
		return &pb.TeamResponse{Success: false, Message: "Failed deleting document"}, &NoDocumentError
	} else if !res.Acknowledged {
		return &pb.TeamResponse{Success: false, Message: "Failed deleting document"}, &DBWriteError
	} else if err != nil {
		return &pb.TeamResponse{Success: false, Message: "Failed deleting document"}, err
	}
	return &pb.TeamResponse{Success: true, Id: req.TeamId, Message: "Team deleted successfully"}, nil
}
