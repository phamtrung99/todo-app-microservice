package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phamtrung99/todo-app-microservice/api-gateway/client"
	"github.com/phamtrung99/todo-app-microservice/api-gateway/config"
	"github.com/phamtrung99/todo-app-microservice/api-gateway/handler"
	"go-micro.dev/v5"
	"go-micro.dev/v5/logger"
)

func main() {
	// Load config
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	cfg := config.Get()

	// Initialize Go-Micro Service
	service := micro.NewService(micro.Name(cfg.ServiceName))
	service.Init()

	// Create gRPC Client
	todoClient := client.NewTodoClient(service.Client())

	// Create HTTP Router
	r := mux.NewRouter()
	todoHandler := handler.NewTodoHandler(todoClient)

	// Define API Routes
	r.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	r.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")

	// Start HTTP server
	logger.Info("API Gateway running on port 3000...")
	http.ListenAndServe(":3000", r)
}
