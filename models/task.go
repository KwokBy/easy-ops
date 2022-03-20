package models

type Task struct {
	ID       int64  `gorm:"column:id" json:"id" form:"id"`
	Username string `gorm:"column:username" json:"username" form:"username"`
	Name     string `gorm:"column:name" json:"name" form:"name"`
	Host     string `gorm:"column:host" json:"host" form:"host"`
	Content  string `gorm:"column:content" json:"content" form:"content"`
	FileType string `gorm:"column:file_type" json:"file_type" form:"file_type"`
	HostIDs  string `gorm:"column:host_ids" json:"host_ids" form:"host_ids"`
	Spec     string `gorm:"column:spec" json:"spec" form:"spec"`
	ExecIds  string `gorm:"column:exec_ids" json:"exec_ids" form:"exec_ids"`
}

func (t *Task) TableName() string {
	return "t_task"
}
