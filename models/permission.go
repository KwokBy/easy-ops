package models

type Permission struct {
	RoleID  int64   `gorm:"column:role_id" json:"role_id" form:"role_id"`
	MenuIDs []int64 `gorm:"column:menu_ids" json:"menu_ids" form:"menu_ids"`
	ApiIDs  []int64 `gorm:"column:api_ids" json:"api_ids" form:"api_ids"`
}
