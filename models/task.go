package models

import (
	"time"

	"github.com/KwokBy/easy-ops/pkg/str"
)

type Task struct {
	ID          int64     `gorm:"primary_key" json:"id" form:"id"`
	Username    string    `gorm:"column:username" json:"username" form:"username"`
	Name        string    `gorm:"column:name" json:"name" form:"name"`
	Content     string    `gorm:"column:content" json:"content" form:"content"`
	FileType    string    `gorm:"column:file_type" json:"file_type" form:"file_type"`
	HostIDs     string    `gorm:"column:host_ids" json:"host_ids" form:"host_ids"`
	Spec        string    `gorm:"column:spec" json:"spec" form:"spec"`
	ExecIds     string    `gorm:"column:exec_ids" json:"exec_ids" form:"exec_ids"`
	Status      int       `gorm:"column:status" json:"status" form:"status"`
	Desc        string    `gorm:"column:desc" json:"desc" form:"desc"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time" form:"updated_time"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time" form:"created_time"`
}

func (t *Task) TableName() string {
	return "t_task"
}

func (t *Task) ToDTO() (TaskDTO, error) {
	hostIDs, err := str.String2Int64s(t.HostIDs)
	if err != nil {
		return TaskDTO{}, err
	}
	execIDs, err := str.String2Int64s(t.ExecIds)
	if err != nil {
		return TaskDTO{}, err
	}
	return TaskDTO{
		ID:          t.ID,
		Username:    t.Username,
		Name:        t.Name,
		Content:     t.Content,
		FileType:    t.FileType,
		HostIDs:     hostIDs,
		Spec:        t.Spec,
		ExecIds:     execIDs,
		Status:      t.Status,
		Desc:        t.Desc,
		UpdatedTime: t.UpdatedTime,
		CreatedTime: t.CreatedTime,
	}, err
}
