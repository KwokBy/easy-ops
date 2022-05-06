package models

import "time"

type Menu struct {
	ID          int64     `gorm:"column:id" json:"id"`
	MenuID      int64     `gorm:"column:menu_id" json:"menu_id"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"`
	ParentId    string    `gorm:"column:parent_id" json:"parent_id"` //父菜单ID
	Path        string    `gorm:"column:path" json:"path"`           //路由path
	Name        string    `gorm:"column:name" json:"name"`           //路由name
	Component   string    `gorm:"column:component" json:"component"` //对应前端文件路径
	Redirect    string    `gorm:"column:redirect" json:"redirect"`   //路由重定向
	Meta        Meta      `gorm:"comment:附加属性" json: "meta"`         //路由meta
}
type Meta struct {
	Title           string `json:"title" gorm:"comment:菜单名"`              // 菜单名
	Icon            string `json:"icon" gorm:"comment:菜单图标"`              // 菜单图标
	KeepAlive       bool   `json:"keepAlive" gorm:"comment:是否缓存"`         // 是否缓存
	ShowLink        bool   `json:"showLink" gorm:"comment:是否在菜单中显示"`      // 是否在菜单中显示
	RefreshRedirect string `json:"refreshRedirect" gorm:"comment:刷新后重定向"` // 刷新后重定向
	Rank            int    `json:"rank" gorm:"comment:排序"`                // 排序
}
