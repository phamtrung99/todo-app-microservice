package main

import (
	"context"
	"fmt"

	pb "github.com/phamtrung99/todo-app-microservice/todo/proto"
	"go-micro.dev/v5"
	"go-micro.dev/v5/logger"
)

func main() {
	service := micro.NewService(micro.Name("todo.service"))

	service.Init()

	todoService := &TodoService{
		todos: make(map[string]*pb.Todo),
	}
	pb.RegisterTodoServiceHandler(service.Server(), todoService)

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

type TodoService struct {
	todos map[string]*pb.Todo // TODO: need db to storage this
}

func (s *TodoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest, rsp *pb.CreateTodoResponse) error {
	id := fmt.Sprintf("%d", len(s.todos)+1)
	s.todos[id] = &pb.Todo{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
	}
	logger.Infof("Created todo with ID: %s", id)

	rsp.Id = id
	return nil
}

func (s *TodoService) GetTodos(ctx context.Context, req *pb.GetTodosRequest, rsp *pb.GetTodosResponse) error {
	for _, todo := range s.todos {
		rsp.Todos = append(rsp.Todos, todo)
	}

	logger.Infof("Get all todo")

	return nil
}
