package repo

import (
	pb "collaboration/pb/team_pb"
)

func ValidateMembers(members []*pb.Member) error {
	for _, member := range members {
		// TODO: check if the user exists: user service is not ready yet
		if member.UserId == 0 {
			return &BadUserError{}
		} 
	}
	return nil
}