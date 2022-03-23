package models

import "time"

type ExecHistoryInfo struct {
	ID          int64     `gorm:"column:id" json:"id"`
	TaskId      int64     `gorm:"column:task_id" json:"task_id"`
	Type        int       `gorm:"column:type" json:"type"`     //执行类型，测试（0），定时任务（1）
	Status      int       `gorm:"column:status" json:"status"` //成功（1）,失败（0）
	Content     string    `gorm:"column:content" json:"content"`
	HostName    string    `gorm:"column:host_name" json:"host_name"` //主机名
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"`
	ExecId      int       `gorm:"column:exec_id" json:"exec_id"` //执行id用来标识同一批执行的任务
	TimeConsume float64   `gorm:"column:time_consume" json:"time_consume"` // 耗时
}

func (e *ExecHistoryInfo) TableName() string {
	return "t_exec_history_info"
}
