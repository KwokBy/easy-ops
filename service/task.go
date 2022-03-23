package service

import (
	"context"
	"time"

	"github.com/KwokBy/easy-ops/models"
	"github.com/KwokBy/easy-ops/pkg/ssh"
	"github.com/KwokBy/easy-ops/pkg/str"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/KwokBy/easy-ops/repo"
	"github.com/robfig/cron/v3"
)

type taskService struct {
	taskRepo            repo.TaskRepo
	cron                *cron.Cron
	execHistoryInfoRepo repo.ExecHistoryInfoRepo
	execHistoryRepo     repo.ExecHistoryRepo
}

func NewTaskService(taskRepo repo.TaskRepo,
	execHistoryInfoRepo repo.ExecHistoryInfoRepo,
	execHistoryRepo repo.ExecHistoryRepo) TaskService {
	return &taskService{
		taskRepo:            taskRepo,
		cron:                cron.New(),
		execHistoryInfoRepo: execHistoryInfoRepo,
		execHistoryRepo:     execHistoryRepo,
	}
}

// AddRunCmdTask 添加运行命令任务
func (s *taskService) AddTask(ctx context.Context, taskDTO models.TaskDTO) error {
	// 数据加入数据库
	task, err := taskDTO.ToPOJO()
	if err != nil {
		return err
	}
	task.CreatedTime = time.Now()
	task.UpdatedTime = time.Now()
	task.Status = TaskStatusInactivated
	task.Username = "doubleguo"

	s.taskRepo.AddTask(ctx, task)
	return nil
}

// AddRunCmdTask 添加运行命令任务
func (s *taskService) AddTaskAndRun(ctx context.Context, taskDTO models.TaskDTO) error {
	// 数据加入数据库
	taskDTO.CreatedTime = time.Now()
	taskDTO.UpdatedTime = time.Now()
	task, err := taskDTO.ToPOJO()
	if err != nil {
		return err
	}
	if err := s.taskRepo.AddTask(ctx, task); err != nil {
		zlog.Errorf("add task error, err: %v", err)
		return err
	}
	if err := s.ExecuteTask(ctx, taskDTO); err != nil {
		zlog.Errorf("execute task error, err: %v", err)
		return err
	}
	return nil
}

// ExecuteTask 执行任务
func (s *taskService) ExecuteTask(ctx context.Context, taskDTO models.TaskDTO) error {

	// 获取主机列表
	_, hosts, err := s.taskRepo.GetTaskAndHosts(ctx, taskDTO.ID, taskDTO.HostIDs)
	if err != nil {
		zlog.Errorf("get task and hosts error, err: %v", err)
		return err
	}
	var entryIDs []int64
	// 执行命令
	for _, host := range hosts {
		entryID, err := s.cron.AddFunc(taskDTO.Spec, func() {
			ssh.ClientAndExec(host, taskDTO.Content)
		})
		if err != nil {
			zlog.Errorf("host %s add task , err: %v", host.Name, err)
			return err
		}
		entryIDs = append(entryIDs, int64(entryID))
	}
	taskDTO.ExecIds = entryIDs
	task, err := taskDTO.ToPOJO()
	if err != nil {
		return err
	}
	task.Status = TaskStatusNotSchedule
	if err := s.taskRepo.UpdateTask(ctx, task); err != nil {
		zlog.Errorf("update task error, err: %v", err)
		return err
	}
	s.cron.Start()
	return nil
}

const (
	ExecHistoryTypeTest      = 1
	ExecHistoryTypeRun       = 2
	ExecHistoryStatusSuccess = 1
	ExecHistoryStatusFailed  = 0
)

// ExecuteTestTask 执行测试任务
func (s *taskService) ExecuteTest(ctx context.Context, taskDTO models.TaskDTO) error {
	// 获取主机列表
	_, hosts, err := s.taskRepo.GetTaskAndHosts(ctx, taskDTO.ID, taskDTO.HostIDs)
	if err != nil {
		zlog.Errorf("get task and hosts error, err: %v", err)
		return err
	}
	var execHistorys []models.ExecHistoryInfo
	// 计算执行id，对于每个任务唯一
	execID, err := s.execHistoryRepo.GetExecHistoryCountByTaskID(ctx, taskDTO.ID)
	if err != nil {
		zlog.Errorf("get count group by exec id error, err: %v", err)
		return err
	}
	execID++
	execHistory := models.ExecHistory{
		TaskID:   taskDTO.ID,
		ExecID:   execID,
		ExecTime: time.Now(),
		Status:   ExecHistoryStatusSuccess,
	}
	for _, host := range hosts {
		currentTime := time.Now()
		execHistoryInfo := models.ExecHistoryInfo{
			TaskId:      taskDTO.ID,
			HostName:    host.Name,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
			Status:      ExecHistoryStatusSuccess,
			Type:        ExecHistoryTypeTest,
			ExecId:      execID,
		}
		// 执行命令
		result, err := ssh.ClientAndExec(host, taskDTO.Content)
		timeConsume := time.Since(currentTime)
		if err != nil {
			zlog.Errorf("host %s exec task , err: %v", host.Name, err)
			execHistoryInfo.Status = ExecHistoryStatusFailed
			execHistory.Status = ExecHistoryStatusFailed
		}
		execHistoryInfo.Content = result
		execHistoryInfo.TimeConsume = timeConsume.Seconds()
		execHistorys = append(execHistorys, execHistoryInfo)
	}
	// 保存执行记录
	if err := s.execHistoryRepo.AddExecHistory(ctx, execHistory); err != nil {
		zlog.Errorf("add exec history error, err: %v", err)
		return err
	}
	if err := s.execHistoryInfoRepo.BatchAddExecHistory(ctx, execHistorys); err != nil {
		zlog.Errorf("batch add exec history info error, err: %v", err)
		return err
	}
	return nil
}

const (
	// TaskStatusInactivated 未激活
	TaskStatusInactivated = 0
	// TaskStatusNotSchedule 待调度
	TaskStatusNotSchedule = 1
	// TaskStatusRun 执行中
	TaskStatusRun = 2
	// TaskStatusFinish 已完成
	TaskStatusFinish = 3
	// TaskStatusFail 失败
	TaskStatusFail = 4
	// TaskStatusCancel 已取消
	TaskStatusCancel = 5
)

// 停止某个任务，
func (s *taskService) StopTask(ctx context.Context, id int64) error {
	task, err := s.taskRepo.GetTaskByID(ctx, id)
	if err != nil {
		zlog.Errorf("get task by id error, err: %v", err)
		return err
	}
	s.cron.Stop()
	zlog.Infof("stop task %s", task)
	entryIDs, err := str.String2Int64s(task.ExecIds)
	if err != nil {
		zlog.Errorf("task.exec_ids to int64 array error, err: %v", err)
		return err
	}
	for _, entryID := range entryIDs {
		s.cron.Remove(cron.EntryID(entryID))
	}
	if err := s.taskRepo.UpdateTaskStatus(ctx, id, TaskStatusInactivated); err != nil {
		zlog.Errorf("update task status error, err: %v", err)
		return err
	}
	s.cron.Start()
	return nil
}

// DeleteTask 删除任务
func (s *taskService) DeleteTask(ctx context.Context, id int64) error {

	return nil
}

// UpdateTask 更新任务
func (s *taskService) UpdateTask(ctx context.Context, task models.Task) error {

	return nil
}

// GetTasksByUsername 获取用户的任务列表
func (s *taskService) GetTasksByUsername(ctx context.Context, userName string) ([]models.TaskDTO, error) {
	tasks, err := s.taskRepo.GetTasksByUsername(ctx, userName)
	if err != nil {
		zlog.Errorf("get tasks by user name error, err: %v", err)
		return nil, err
	}
	var taskDTOs []models.TaskDTO
	for _, task := range tasks {
		taskDTO, err := task.ToDTO()
		if err != nil {
			zlog.Errorf("task to dto error, err: %v", err)
			return nil, err
		}
		taskDTOs = append(taskDTOs, taskDTO)
	}
	return taskDTOs, nil
}
