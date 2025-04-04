package main

import (
	"github.com/phamtrung99/todo-app-microservice/todo/config"
	pb "github.com/phamtrung99/todo-app-microservice/todo/proto"
	tdService "github.com/phamtrung99/todo-app-microservice/todo/service"
	"github.com/phamtrung99/todo-app-microservice/todo/utils"
	"go-micro.dev/v5"
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/server"
)

func main() {
	// Load config
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	cfg := config.Get()

	// Init service
	service := micro.NewService(micro.Name(cfg.ServiceName))

	service.Server().Init(
		server.Wait(nil), // graceful shutdown
	)

	// Init DB connect
	dbClient := utils.GetMysqlClient(cfg)
	err = utils.Migrate(cfg)
	if err != nil {
		panic(err)
	}

	todoService := tdService.NewTodoService(dbClient)

	pb.RegisterTodoServiceHandler(service.Server(), todoService)

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
