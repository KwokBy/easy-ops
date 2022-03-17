package models

import "time"

type Host struct {
	ID          int64     `gorm:"column:id" json:"id" form:"id"`
	Owner       string    `gorm:"column:owener" json:"owner" form:"owner"`
	HostName    string    `gorm:"column:host_name" json:"host_name" form:"host_name"`
	Host        string    `gorm:"column:host" json:"host" form:"host"`
	Name        string    `gorm:"column:name" json:"name" form:"name"`
	Desc        string    `gorm:"column:desc" json:"desc" form:"desc"`
	Port        int64     `gorm:"column:port" json:"port" form:"port"`
	Password    string    `gorm:"column:password" json:"password" form:"password"`
	SSHType     string    `gorm:"column:ssh_type" json:"ssh_type" form:"ssh_type"`
	SSHKeyPath  string    `gorm:"column:ssh_key_path" json:"ssh_key_path" form:"ssh_key_path"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time" form:"updated_time"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time" form:"created_time"`
}

func (h *Host) TableName() string {
	return "t_host"
}
