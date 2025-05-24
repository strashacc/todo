package repo

import (
	pb "collaboration/pb/team_pb"
	// user_pb "collaboration/pb/user/generated"
	"collaboration/clients/user_client"
)

func ValidateMembers(members []*pb.Member) error {
	for _, member := range members {
		user, err := userclient.GetUser(member.UserId)
		if user.Id == 0 {
			return &BadUserError
		} else if err != nil {
			return &BadUserError
		}
	}
	return nil
}