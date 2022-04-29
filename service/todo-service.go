package service

import (
	context "context"

	"github.com/satioO/todo-grpc/api"
)

type todoService struct {
	api.UnimplementedTodoServiceServer
}

func NewTodoService() api.TodoServiceServer {
	return &todoService{}
}

// GetTodos implements api.TodoServiceServer
func (*todoService) GetTodos(context.Context, *api.EmptyRequest) (*api.TodoResponse, error) {
	return &api.TodoResponse{
		Id:     1,
		Title:  "Hello, Vaibhav",
		Status: 1,
	}, nil
}
