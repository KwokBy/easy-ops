package repo

import (
	"context"
	"fmt"

	"github.com/KwokBy/easy-ops/models"
	"gorm.io/gorm"
)

type mysqlMirrorRepo struct {
	DB *gorm.DB
}

func NewMysqlMirrorRepo(DB *gorm.DB) IMirrorRepo {
	return &mysqlMirrorRepo{DB}
}

// GetMirrorsByAdmin 根据管理员获取镜像列表
func (m *mysqlMirrorRepo) GetMirrorsByAdmin(ctx context.Context, admin string) (
	[]models.Mirror, error) {
	var mirrors []models.Mirror
	if err := m.DB.Raw("SELECT a.* FROM t_mirror a INNER JOIN (SELECT name, MAX(version) version FROM t_mirror GROUP BY `name`) b ON a.name = b.name and a.version = b.version and LOCATE(?,`admin`) order by a.name;", admin).
		Scan(&mirrors).Error; err != nil {
		return nil, err
	}
	return mirrors, nil
}

// AddAdmin 添加管理员
func (m *mysqlMirrorRepo) AddAdmin(ctx context.Context, mirrorId, admin string) error {
	var mirror models.Mirror
	if err := m.DB.Model(&mirror).Where("id = ?", mirrorId).Error; err != nil {
		return err
	} else {
		mirror.Admin = fmt.Sprintf("%s,%s", mirror.Admin, admin)
		if err := m.DB.Save(&mirror).Error; err != nil {
			return err
		}
	}
	return nil
}

// AddMirror 添加镜像
func (m *mysqlMirrorRepo) AddMirror(ctx context.Context, mirror models.Mirror) error {
	if err := m.DB.Create(&mirror).Error; err != nil {
		return err
	}
	return nil
}

// DeleteMirror 删除镜像
func (m *mysqlMirrorRepo) DeleteMirror(ctx context.Context, id int64) error {
	if err := m.DB.Delete(&models.Mirror{}, id).Error; err != nil {
		return err
	}
	return nil
}
