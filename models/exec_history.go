package models

import "time"

type ExecHistory struct {
	Id       int64     `gorm:"column:id" json:"id"`
	ExecID   int64       `gorm:"column:exec_id" json:"exec_id"`
	ExecTime time.Time `gorm:"column:exec_time" json:"exec_time"`
	TaskID   int64     `gorm:"column:task_id" json:"task_id"`
	Status   int64     `gorm:"column:status" json:"status"` //一组任务只有有一个失败视为失败
}

func (e *ExecHistory) TableName() string {
	return "t_exec_history"
}
