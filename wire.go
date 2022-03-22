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

// router 解决参数过多
var router = wire.NewSet(wire.Struct(new(api.Router), "*"))

// InitServer Injectors from wire.go:
func InitServer() *app.Server {
	wire.Build(
		app.InitGormMySql,
		repo.NewMysqlDemoRepo,
		repo.NewMysqlHostRepo,
		repo.NewMysqlTaskRepo,
		repo.NewMysqlExecHistoryRepo,
		service.NewDemoService,
		service.NewHostService,
		service.NewTaskService,
		service.NewExecHistoryService,
		handlers.NewDemoHandler,
		handlers.NewUserHandler,
		handlers.NewWsSshHandler,
		handlers.NewHostHandler,
		handlers.NewTaskHandler,
		handlers.NewExecHistoryHandler,
		router,
		app.NewServer,
		app.NewGinEngine,
	)
	return &app.Server{}
}
