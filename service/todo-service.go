package service

import (
	"context"

	"github.com/satioO/todo-grpc/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoService struct {
	api.UnimplementedTodoServiceServer
	db *mongo.Database
}

func NewTodoService(db *mongo.Database) api.TodoServiceServer {
	return &todoService{
		db: db,
	}
}

// GetTodos implements api.TodoServiceServer
func (t *todoService) GetTodos(ctx context.Context, _ *api.EmptyRequest) (*api.TodoResponse, error) {
	todos := []*api.Todo{}

	cur, err := t.db.Collection("todos").Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		todo := api.Todo{}
		err := cur.Decode(&todo)
		if err != nil {
			return nil, err
		}

		todos = append(todos, &todo)
	}

	cur.Close(context.TODO())

	if len(todos) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return &api.TodoResponse{
		Items: todos,
	}, nil
}

// CreateTodo implements api.TodoServiceServer
func (t *todoService) CreateTodo(ctx context.Context, todo *api.Todo) (*api.CreateTodoResponse, error) {
	_, err := t.db.Collection("todos").InsertOne(context.TODO(), todo)

	if err != nil {
		return nil, err
	}

	return &api.CreateTodoResponse{
		Message: "Todo Created :::",
	}, nil
}
