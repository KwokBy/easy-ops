// Package repo 数据层操作
package repo

import (
	"context"
	"github.com/KwokBy/easy-ops/models"
)

// IDemoRepo represent the demo repository contract
type IDemoRepo interface {
	// GetDemos return all demos
	GetDemos(ctx context.Context) ([]models.Demo, error)
}
