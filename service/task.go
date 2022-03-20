package service

import (
	"context"
	"strings"

	"github.com/KwokBy/easy-ops/models"
	"github.com/KwokBy/easy-ops/pkg/ssh"
	"github.com/KwokBy/easy-ops/pkg/str"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/KwokBy/easy-ops/repo"
	"github.com/robfig/cron/v3"
)

type taskService struct {
	taskRepo repo.TaskRepo
	cron     *cron.Cron
}

func NewTaskService(taskRepo repo.TaskRepo) TaskService {
	return &taskService{
		taskRepo: taskRepo,
		cron:     cron.New(),
	}
}

// AddRunCmdTask 添加运行命令任务
func (s *taskService) AddTask(ctx context.Context, task models.Task) error {
	// 数据加入数据库
	s.taskRepo.AddTask(ctx, task)
	return nil
}

// AddRunCmdTask 添加运行命令任务
func (s *taskService) AddTaskAndRun(ctx context.Context, task models.Task) error {
	// 数据加入数据库
	if err := s.taskRepo.AddTask(ctx, task); err != nil {
		zlog.Errorf("add task error, err: %v", err)
		return err
	}
	if err := s.ExecuteTask(ctx, task); err != nil {
		zlog.Errorf("execute task error, err: %v", err)
		return err
	}
	return nil
}

// ExecuteTask 执行任务
func (s *taskService) ExecuteTask(ctx context.Context, task models.Task) error {

	hostIDs, err := str.Strings2Int64s(strings.Split(task.HostIDs, ","))
	if err != nil {
		zlog.Errorf("task.host_ids to int64 array error, err: %v", err)
		return err
	}
	// 获取主机列表
	_, hosts, err := s.taskRepo.GetTaskAndHosts(ctx, task.ID, hostIDs)
	if err != nil {
		zlog.Errorf("get task and hosts error, err: %v", err)
		return err
	}
	var entryIDs []int64
	// 执行命令
	for _, host := range hosts {
		entryID, err := s.cron.AddFunc(task.Spec, func() {
			ssh.ClientAndExec(host, task.Content)
		})
		if err != nil {
			zlog.Errorf("host %s add task , err: %v", host.Name, err)
			return err
		}
		entryIDs = append(entryIDs, int64(entryID))
	}
	task.ExecIds, err = str.Int64s2String(entryIDs)
	if err != nil {
		zlog.Errorf("entryIDs to string error, err: %v", err)
		return err
	}
	if err := s.taskRepo.UpdateTaskEntryId(ctx, task.ID, task.ExecIds); err != nil {
		zlog.Errorf("update task entry id error, err: %v", err)
		return err
	}
	s.cron.Start()
	return nil
}

const (
	// TaskStatusRun 执行中
	TaskStatusRun = 1
	// TaskStatusStop 停止
	TaskStatusStop = 2
)

// 停止某个任务，
func (s *taskService) StopTask(ctx context.Context, id int64) error {
	task, err := s.taskRepo.GetTaskByID(ctx, id)
	if err != nil {
		zlog.Errorf("get task by id error, err: %v", err)
		return err
	}
	s.cron.Stop()
	entryIDs, err := str.Strings2Int64s(strings.Split(task.ExecIds, ","))
	if err != nil {
		zlog.Errorf("task.exec_ids to int64 array error, err: %v", err)
		return err
	}
	for _, entryID := range entryIDs {
		s.cron.Remove(cron.EntryID(entryID))
	}
	if err := s.taskRepo.UpdateTaskStatus(ctx, id, TaskStatusStop); err != nil {
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
