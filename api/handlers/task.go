package handlers

import (
	"github.com/KwokBy/easy-ops/models"
	"github.com/KwokBy/easy-ops/pkg/response"
	"github.com/KwokBy/easy-ops/pkg/zlog"
	"github.com/KwokBy/easy-ops/service"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) TaskHandler {
	return TaskHandler{
		taskService: taskService,
	}
}

type GetTasksReq struct {
	Username string `json:"owner"`
}

// GetTasks 获取任务列表
func (h *TaskHandler) GetTasks(c *gin.Context) {
	var req GetTasksReq
	if err := c.ShouldBind(&req); err != nil {
		zlog.Errorf("get owner form uri error: %s", err.Error())
		response.FailWithData(err, "get owner form uri error", c)
		return
	}
	tasks, err := h.taskService.GetTasksByUsername(c, req.Username)
	if err != nil {
		zlog.Errorf("get tasks by owner error: %s", err.Error())
		response.FailWithData(err, "get tasks by owner error", c)
		return
	}
	response.OKWithData(tasks, "get tasks by owner success", c)
}

// AddTask 添加任务
func (h *TaskHandler) AddTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBind(&task); err != nil {
		zlog.Errorf("add task error: %s", err.Error())
		response.FailWithData(err, "add task error", c)
		return
	}
	if err := h.taskService.AddTask(c, task); err != nil {
		zlog.Errorf("add task error: %s", err.Error())
		response.FailWithData(err, "add task error", c)
		return
	}
	response.OKWithData(nil, "add task success", c)
}

// AddTaskAndExecute 添加任务并执行
func (h *TaskHandler) AddTaskAndExecute(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBind(&task); err != nil {
		zlog.Errorf("add task error: %s", err.Error())
		response.FailWithData(err, "add task error", c)
		return
	}
	if err := h.taskService.AddTaskAndRun(c, task); err != nil {
		zlog.Errorf("add task error: %s", err.Error())
		response.FailWithData(err, "add task error", c)
		return
	}
	response.OKWithData(nil, "add task success", c)
}

// ExecuteTask 执行任务
func (h *TaskHandler) ExecuteTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBind(&task); err != nil {
		zlog.Errorf("execute task error: %s", err.Error())
		response.FailWithData(err, "execute task error", c)
		return
	}
	if err := h.taskService.ExecuteTask(c, task); err != nil {
		zlog.Errorf("execute task error: %s", err.Error())
		response.FailWithData(err, "execute task error", c)
		return
	}
	response.OKWithData(nil, "execute task success", c)
}

// StopTasks 停止任务 struct
type StopTasksReq struct {
	TaskID int64 `json:"id"`
}

// StopTask 停止任务
func (h *TaskHandler) StopTask(c *gin.Context) {
	var req StopTasksReq
	if err := c.ShouldBind(&req); err != nil {
		zlog.Errorf("stop task error: %s", err.Error())
		response.FailWithData(err, "stop task error", c)
		return
	}
	if err := h.taskService.StopTask(c, req.TaskID); err != nil {
		zlog.Errorf("stop task error: %s", err.Error())
		response.FailWithData(err, "stop task error", c)
		return
	}
	response.OKWithData(nil, "stop task success", c)
}
