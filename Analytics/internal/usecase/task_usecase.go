package usecase

import "context"

type Stats struct {
	CompletedToday int32
	CompletedWeek  int32
	TotalTasks     int32
}

type TaskUsecase interface {
	GetStats() (*Stats, error)
}

type taskUsecase struct {
	repo interface {
		CountStats(ctx context.Context) (int64, int64, int64, error) // <-- здесь int64
	}
}

func NewTaskUsecase(r interface {
	CountStats(ctx context.Context) (int64, int64, int64, error) // <-- и здесь int64
}) TaskUsecase {
	return &taskUsecase{repo: r}
}

func (u *taskUsecase) GetStats() (*Stats, error) {
	ctx := context.TODO()
	today64, week64, total64, err := u.repo.CountStats(ctx)
	if err != nil {
		return nil, err
	}
	return &Stats{
		CompletedToday: int32(today64),
		CompletedWeek:  int32(week64),
		TotalTasks:     int32(total64),
	}, nil
}
