//go:build wireinject
// +build wireinject

package main

import (
	"github.com/KwokBy/easy-ops/api"
	"github.com/KwokBy/easy-ops/api/handlers"
	"github.com/KwokBy/easy-ops/app"
	"github.com/KwokBy/easy-ops/repo"
	"github.com/KwokBy/easy-ops/service"
	"github.com/google/wire"
)

// InitServer Injectors from wire.go:
func InitServer() *app.Server {
	wire.Build(
		app.InitGormMySql,
		repo.NewMysqlDemoRepo,
		service.NewDemoService,
		handlers.NewDemoHandler,
		api.NewRouter,
		app.NewServer,
		app.NewGinEngine,
	)
	return &app.Server{}
}
