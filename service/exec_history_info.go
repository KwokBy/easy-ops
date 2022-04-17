package service

import (
	"context"

	"github.com/KwokBy/easy-ops/models"
	"github.com/KwokBy/easy-ops/repo"
)

type execHistoryInfoService struct {
	execHistoryInfoRepo repo.ExecHistoryInfoRepo
}

func NewExecHistoryInfoService(execHistoryInfoRepo repo.ExecHistoryInfoRepo) ExecHistoryInfoService {
	return &execHistoryInfoService{
		execHistoryInfoRepo: execHistoryInfoRepo,
	}
}

// GetExecHistoryInfos 获取执行历史列表
func (s *execHistoryInfoService) GetExecHistoryInfos(ctx context.Context, taskID, execID int64) (models.ExecHistoryDTO, error) {
	execHistoryInfos, err := s.execHistoryInfoRepo.GetExecHistoryByTaskIDAndExecID(ctx, taskID, execID)
	if err != nil {
		return models.ExecHistoryDTO{}, err
	}
	resultCount := [2]int{}
	sumCount := 0.0
	for _, execHistoryInfo := range execHistoryInfos {
		resultCount[execHistoryInfo.Status]++
		sumCount += execHistoryInfo.TimeConsume
	}
	return models.ExecHistoryDTO{
		ExecHistories: execHistoryInfos,
		SuccessCount:  resultCount[1],
		FailCount:     resultCount[0],
		AvgCount:      sumCount / float64(len(execHistoryInfos)),
	}, nil
}
