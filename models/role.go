package models

import "time"

type Role struct {
	ID          int64     `gorm:"column:id" json:"id" form:"id"`
	RoleId      int64     `gorm:"column:role_id" json:"role_id" form:"role_id"`
	NameZh      string    `gorm:"column:name_zh" json:"name_zh" form:"name_zh"`
	NameEn      string    `gorm:"column:name_en" json:"name_en" form:"name_en"`
	Desc        string    `gorm:"column:desc" json:"desc" form:"desc"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time" form:"updated_time"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time" form:"created_time"`
}

func (r *Role) TableName() string {
	return "t_role"
}
