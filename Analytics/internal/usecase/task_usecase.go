package usecase

import "context"

type Stats struct {
	TotalTasks     int32
	CompletedTasks int32
	InProgress     int32
	Overdue        int32
}

type TaskUsecase interface {
	GetStats() (*Stats, error)
	GetTeamStats(teamID string) (*Stats, error)
	GetUserStats(userID string) (*Stats, error)
}

type taskUsecase struct {
	repo interface {
		CountStats(ctx context.Context) (total, completed, inProgress, overdue int64, err error)
		CountTeamStats(ctx context.Context, teamID string) (total, completed, inProgress, overdue int64, err error)
		CountUserStats(ctx context.Context, userID string) (total, completed, inProgress, overdue int64, err error)
	}
}

func NewTaskUsecase(r interface {
	CountStats(ctx context.Context) (total, completed, inProgress, overdue int64, err error)
	CountTeamStats(ctx context.Context, teamID string) (total, completed, inProgress, overdue int64, err error)
	CountUserStats(ctx context.Context, userID string) (total, completed, inProgress, overdue int64, err error)
}) TaskUsecase {
	return &taskUsecase{repo: r}
}

func (u *taskUsecase) GetStats() (*Stats, error) {
	return convertStats(u.repo.CountStats(context.TODO()))
}

func (u *taskUsecase) GetTeamStats(teamID string) (*Stats, error) {
	return convertStats(u.repo.CountTeamStats(context.TODO(), teamID))
}

func (u *taskUsecase) GetUserStats(userID string) (*Stats, error) {
	return convertStats(u.repo.CountUserStats(context.TODO(), userID))
}

func convertStats(total, completed, inProgress, overdue int64, err error) (*Stats, error) {
	if err != nil {
		return nil, err
	}
	return &Stats{
		TotalTasks:     int32(total),
		CompletedTasks: int32(completed),
		InProgress:     int32(inProgress),
		Overdue:        int32(overdue),
	}, nil
}
