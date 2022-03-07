package models

type Task struct {
	ID       int64  `gorm:"column:id" json:"id" form:"id"`
	Username string `gorm:"column:username" json:"username" form:"username"`
	Host     string `gorm:"column:host" json:"host" form:"host"`
	Content  string `gorm:"column:content" json:"content" form:"content"`
	FileType string `gorm:"column:file_type" json:"file_type" form:"file_type"`
}

func (t *Task) TableName() string {
	return "t_task"
}
