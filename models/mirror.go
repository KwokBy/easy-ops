package models

import "time"

type Mirror struct {
	ID          int64     `gorm:"column:id" json:"id" form:"id"`
	MirrorId    string    `gorm:"column:mirror_id" json:"mirror_id" form:"mirror_id"`
	Creator     string    `gorm:"column:creator" json:"creator" form:"creator"`
	Name        string    `gorm:"column:name" json:"name" form:"name"`
	Version     string    `gorm:"column:version" json:"version" form:"version"`
	Tag         string    `gorm:"column:tag" json:"tag" form:"tag"`
	Desc        string    `gorm:"column:desc" json:"desc" form:"desc"`
	Admin       string    `gorm:"column:admin" json:"admin" form:"admin"`
	Content     string    `gorm:"column:content" json:"content" form:"content"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time" form:"updated_time"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time" form:"created_time"`
}

func (m *Mirror) TableName() string {
	return "t_mirror"
}
