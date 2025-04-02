package main

import (
	pb "github.com/phamtrung99/todo-app-microservice/todo/proto"
	tdService "github.com/phamtrung99/todo-app-microservice/todo/service"
	"github.com/phamtrung99/todo-app-microservice/todo/utils"
	"go-micro.dev/v5"
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/server"
)

func main() {
	service := micro.NewService(micro.Name("todo.service"))

	service.Server().Init(
		server.Wait(nil), // graceful shutdown
	)

	dbClient := utils.GetMysqlClient()
	err := utils.Migrate()
	if err != nil {
		panic(err)
	}

	todoService := tdService.NewTodoService(dbClient)

	pb.RegisterTodoServiceHandler(service.Server(), todoService)

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
