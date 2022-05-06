package models

type Casbin struct {
	ID     uint64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	Ptype  string `gorm:"column:ptype" db:"ptype" json:"ptype" form:"ptype"`
	RoleID string `gorm:"column:v0" db:"v0" json:"v0" form:"v0"`
	Api    string `gorm:"column:v1" db:"v1" json:"v1" form:"v1"`
	Method string `gorm:"column:v2" db:"v2" json:"v2" form:"v2"`
	V3     string `gorm:"column:v3" db:"v3" json:"v3" form:"v3"`
	V4     string `gorm:"column:v4" db:"v4" json:"v4" form:"v4"`
	V5     string `gorm:"column:v5" db:"v5" json:"v5" form:"v5"`
}

func (c *Casbin) TableName() string {
	return "casbin_rule"
}
