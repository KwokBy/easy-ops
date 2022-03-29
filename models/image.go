package models

import "time"

type Image struct {
	Id          int       `gorm:"column:id" db:"id" json:"id" form:"id"`
	Name        string    `gorm:"column:name" db:"name" json:"name" form:"name"`                         //镜像名
	Dockerfile  string    `gorm:"column:dockerfile" db:"dockerfile" json:"dockerfile" form:"dockerfile"` //dockerfile
	Version     string    `gorm:"column:version" db:"version" json:"version" form:"version"`             //版本
	Owner       string    `gorm:"column:owner" db:"owner" json:"owner" form:"owner"`                     //拥有者
	ImageId     string    `gorm:"column:image_id" db:"image_id" json:"image_id" form:"image_id"`         //镜像id
	IsPublic    bool      `gorm:"column:is_public" db:"is_public" json:"is_public" form:"is_public"`     //是否发布
	Desc        string    `gorm:"column:desc" db:"desc" json:"desc" form:"desc"`                         //描述
	UpdatedTime time.Time `gorm:"column:updated_time" db:"updated_time" json:"updated_time" form:"updated_time"`
	CreatedTime time.Time `gorm:"column:created_time" db:"created_time" json:"created_time" form:"created_time"`
}

func (i *Image) TableName() string {
	return "t_image"
}
