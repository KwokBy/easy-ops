package models

type Api struct {
	Id         int    `gorm:"column:id" db:"id" json:"id" form:"id"`
	ModelsName string `gorm:"column:models_name" db:"models_name" json:"models_name" form:"models_name"` //接口所属模块名
	Name       string `gorm:"column:name" db:"name" json:"name" form:"name"`                             //接口名
	Method     string `gorm:"column:method" db:"method" json:"method" form:"method"`                     //请求方法
	Desc       string `gorm:"column:desc" db:"desc" json:"desc" form:"desc"`                             //接口描述
}

func (a *Api) TableName() string {
	return "t_api"
}
