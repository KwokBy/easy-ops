// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/KwokBy/easy-ops/api"
	"github.com/KwokBy/easy-ops/api/handlers"
	"github.com/KwokBy/easy-ops/app"
	"github.com/KwokBy/easy-ops/repo"
	"github.com/KwokBy/easy-ops/service"
	"github.com/google/wire"
)

// Injectors from wire.go:

// InitServer Injectors from wire.go:
func InitServer() *app.Server {
	engine := app.NewGinEngine()
	db := app.InitGormMySql()
	iDemoRepo := repo.NewMysqlDemoRepo(db)
	iDemoService := service.NewDemoService(iDemoRepo)
	demoHandler := handlers.NewDemoHandler(iDemoService)
	userRepo := repo.NewMysqlUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	wsSshHandler := handlers.NewWsSshHandler()
	hostRepo := repo.NewMysqlHostRepo(db)
	hostService := service.NewHostService(hostRepo)
	hostHandler := handlers.NewHostHandler(hostService)
	taskRepo := repo.NewMysqlTaskRepo(db)
	execHistoryInfoRepo := repo.NewMysqlExecHistoryInfoRepo(db)
	execHistoryRepo := repo.NewMysqlExecHistoryRepo(db)
	taskService := service.NewTaskService(taskRepo, execHistoryInfoRepo, execHistoryRepo)
	taskHandler := handlers.NewTaskHandler(taskService)
	execHistoryService := service.NewExecHistoryService(execHistoryRepo)
	execHistoryHandler := handlers.NewExecHistoryHandler(execHistoryService)
	execHistoryInfoService := service.NewExecHistoryInfoService(execHistoryInfoRepo)
	execHistoryInfoHandler := handlers.NewExecHistoryInfoHandler(execHistoryInfoService)
	imageRepo := repo.NewMysqlImageRepo(db)
	imageService := service.NewImageService(imageRepo)
	imageHandler := handlers.NewImageHandler(imageService)
	roleRepo := repo.NewMysqlRoleRepo(db)
	roleService := service.NewRoleService(roleRepo)
	roleHandler := handlers.NewRoleHandler(roleService)
	apiRouter := &api.Router{
		Demo:            demoHandler,
		User:            userHandler,
		WsSsh:           wsSshHandler,
		Host:            hostHandler,
		Task:            taskHandler,
		ExecHistory:     execHistoryHandler,
		ExecHistoryInfo: execHistoryInfoHandler,
		Image:           imageHandler,
		Role:            roleHandler,
	}
	server := app.NewServer(engine, apiRouter)
	return server
}

// wire.go:

// router 解决参数过多
var router = wire.NewSet(wire.Struct(new(api.Router), "*"))
