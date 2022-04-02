package models

import "time"

type User struct {
	ID           int64     `gorm:"column:id" json:"id" form:"id"`
	Username     string    `gorm:"column:username" json:"username" form:"username"`
	Nickname     string    `gorm:"column:nickname" json:"nickname" form:"nickname"`
	PasswordHash string    `gorm:"column:password_hash" json:"password_hash" form:"password_hash"`
	AccessToken  string    `gorm:"column:access_token" json:"access_token" form:"access_token"`
	TokenExpired string    `gorm:"column:token_expired" json:"token_expired" form:"token_expired"`
	Type         string    `gorm:"column:type" json:"type" form:"type"`
	RoleID       int64     `gorm:"column:role_id" json:"role_id" form:"role_id"`
	LastIp       string    `gorm:"column:last_ip" json:"last_ip" form:"last_ip"`
	WxToken      string    `gorm:"column:wx_token" json:"wx_token" form:"wx_token"`
	UpdatedTime  time.Time `gorm:"column:updated_time" json:"updated_time" form:"updated_time"`
	CreatedTime  time.Time `gorm:"column:created_time" json:"created_time" form:"created_time"`
}

func (u *User)TableName() string {
	return "t_user"
}
