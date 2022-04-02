package casbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func Casbin(db *gorm.DB) (*casbin.Enforcer, error) {

	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer("./model.conf", a)
	if err != nil {
		return nil, err
	}
	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return e, nil
}
