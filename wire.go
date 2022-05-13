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
		repo.NewMysqlUserRepo,
		repo.NewMysqlHostRepo,
		repo.NewMysqlTaskRepo,
		repo.NewMysqlExecHistoryInfoRepo,
		repo.NewMysqlExecHistoryRepo,
		repo.NewMysqlImageRepo,
		repo.NewMysqlRoleRepo,
		repo.NewMysqlApiRepo,
		repo.NewMysqlCasbinRepo,
		service.NewRoleService,
		service.NewDemoService,
		service.NewHostService,
		service.NewUserService,
		service.NewTaskService,
		service.NewExecHistoryInfoService,
		service.NewExecHistoryService,
		service.NewImageService,
		handlers.NewRoleHandler,
		handlers.NewDemoHandler,
		handlers.NewUserHandler,
		handlers.NewWsSshHandler,
		handlers.NewHostHandler,
		handlers.NewTaskHandler,
		handlers.NewExecHistoryInfoHandler,
		handlers.NewExecHistoryHandler,
		handlers.NewImageHandler,
		router,
		app.NewServer,
		app.NewGinEngine,
	)
	return &app.Server{}
}
