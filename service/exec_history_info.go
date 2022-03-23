package service

import (
	"context"
	"sort"

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

// GetExecHistoriesByTaskID 根据任务ID获取执行历史列表
func (s *execHistoryInfoService) GetExecHistoriesByTaskID(ctx context.Context, taskID int64) ([][]models.ExecHistoryInfo, error) {
	execHistories, err := s.execHistoryInfoRepo.GetExecHistoryByTaskID(ctx, taskID)
	if err != nil {
		return nil, err
	}
	return splitSlice(execHistories), nil
}

//按某个字段排序
type sortByExecID []models.ExecHistoryInfo

func (s sortByExecID) Len() int           { return len(s) }
func (s sortByExecID) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortByExecID) Less(i, j int) bool { return s[i].ExecId < s[j].ExecId }

//切片分组
func splitSlice(list []models.ExecHistoryInfo) [][]models.ExecHistoryInfo {
	sort.Sort(sortByExecID(list))
	returnData := make([][]models.ExecHistoryInfo, 0)
	i := 0
	var j int
	for {
		if i >= len(list) {
			break
		}
		for j = i + 1; j < len(list) && list[i].ExecId == list[j].ExecId; j++ {
		}

		returnData = append(returnData, list[i:j])
		i = j
	}
	return returnData
}
