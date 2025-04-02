package service

import (
	"context"

	"github.com/phamtrung99/todo-app-microservice/todo/model"
	pb "github.com/phamtrung99/todo-app-microservice/todo/proto"
	"go-micro.dev/v5/logger"
	"gorm.io/gorm"
)

type ITodoService interface {
	CreateTodo(ctx context.Context, req *pb.CreateTodoRequest, rsp *pb.CreateTodoResponse) error
	GetTodos(ctx context.Context, req *pb.GetTodosRequest, rsp *pb.GetTodosResponse) error
}

type todoService struct {
	dbClient *gorm.DB
}

func NewTodoService(dbClient *gorm.DB) ITodoService {
	return &todoService{
		dbClient: dbClient,
	}
}

func (s *todoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest, rsp *pb.CreateTodoResponse) error {
	todo := &model.Todo{
		Title:       req.Title,
		Description: req.Description,
	}
	s.dbClient.Table("todo").Create(todo)

	logger.Infof("Created todo with ID: %s", todo.ID)

	rsp.Id = string(todo.ID)
	return nil
}

func (s *todoService) GetTodos(ctx context.Context, req *pb.GetTodosRequest, rsp *pb.GetTodosResponse) error {
	listTodos := make([]*pb.Todo, 0)
	s.dbClient.Table("todo").Find(listTodos)

	rsp.Todos = listTodos

	logger.Infof("Get all todo")

	return nil
}
