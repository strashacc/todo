package repo

import (
	db "collaboration/db/mongodb"
	pb "collaboration/pb/team_pb"
	"log"

	// user_pb "collaboration/pb/user/generated"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
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
	col := db.GetDatabase().Collection("teams")

	cur, err := col.Find(context.TODO(), bson.M{})
	if err != nil {
		return &pb.TeamList{}, err
	}
	var res pb.TeamList
	cur.All(context.TODO(), &res)
	if len(res.Teams) == 0 {
		return &pb.TeamList{}, &NoDocumentError
	}
	return &res, nil
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

func UpdateTeam(req *pb.UpdateTeamRequest) (*pb.TeamResponse, error) {
	col := db.GetDatabase().Collection("teams")

	res, err := col.ReplaceOne(context.TODO(), pb.GetTeamRequest{TeamId: req.TeamId}, req.Update)
	if err != nil {
		return &pb.TeamResponse{Success: false, Message: "Error"}, nil
	} else if res.ModifiedCount == 0 {
		return &pb.TeamResponse{Success: false, Message: "Team doesn't exist"}, &NoDocumentError
	}
	return &pb.TeamResponse{Success: true, Id: req.TeamId, Message: "Updated team successfully"}, nil
}

func DeleteUser(userId int) (*pb.TeamResponse, error) {
	col := db.GetDatabase().Collection("teams")

	res, err := col.Find(context.TODO(), bson.M{"Members": userId})
	if err != nil {
		return &pb.TeamResponse{Success: false}, nil
	}
	var teams pb.TeamList
	res.All(context.TODO(), &teams.Teams)

	for _, team := range teams.Teams {
		update := &pb.UpdateTeamRequest{
			TeamId: team.Id,
		}
		for i, user := range team.Members {
			if user.UserId == uint64(userId) {
				team.Members = append(team.Members[:i-1], team.Members[i + 1:]...)
				update.Update = team
				break
			}
		}
		res, err := UpdateTeam(update)
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println(res)
		}
	}
	return &pb.TeamResponse{Success: true}, nil
}