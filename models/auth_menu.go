package models

type AuthMenu struct {
	ID     int64 `gorm:"column:id" json:"id" form:"id"`
	RoleID int64 `gorm:"column:role_id" json:"role_id" form:"role_id"`
	MenuID int64 `gorm:"column:menu_id" json:"menu_id" form:"menu_id"`
}

func (a *AuthMenu) TableName() string {
	return "t_auth_menu"
}
