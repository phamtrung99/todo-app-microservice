package client

import (
	"context"

	pbTodo "github.com/phamtrung99/todo-app-microservice/todo/proto"
	"go-micro.dev/v5/client"
	"go-micro.dev/v5/logger"
)

type ITodoClient interface {
	CreateTodo(title, description string) (*pbTodo.CreateTodoResponse, error)
	GetList() (*pbTodo.GetTodosResponse, error)
}

type todoClient struct {
	client pbTodo.TodoService
}

func NewTodoClient(service client.Client) ITodoClient {
	return &todoClient{
		client: pbTodo.NewTodoService("todo.service", service),
	}
}

func (t *todoClient) CreateTodo(title, description string) (*pbTodo.CreateTodoResponse, error) {
	req := &pbTodo.CreateTodoRequest{
		Title:       title,
		Description: description,
	}
	resp, err := t.client.CreateTodo(context.Background(), req)
	if err != nil {
		logger.Errorf("Error calling CreateTodo: %v", err)
		return nil, err
	}
	return resp, nil
}

func (t *todoClient) GetList() (*pbTodo.GetTodosResponse, error) {
	resp, err := t.client.GetTodos(context.Background(), &pbTodo.GetTodosRequest{})
	if err != nil {
		logger.Errorf("Error calling CreateTodo: %v", err)
		return nil, err
	}
	return resp, nil
}
