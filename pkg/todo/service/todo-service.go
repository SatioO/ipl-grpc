package service

import (
	"context"

	"github.com/satioO/todo-grpc/pkg/todo/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type todoService struct {
	pb.UnimplementedTodoServiceServer
	db *mongo.Database
}

func NewTodoService(db *mongo.Database) pb.TodoServiceServer {
	return &todoService{
		db: db,
	}
}

// GetTodos implements pb.TodoServiceServer
func (t *todoService) GetTodos(context.Context, *pb.EmptyRequest) (*pb.TodoResponse, error) {
	todos := []*pb.Todo{}

	cur, err := t.db.Collection("todos").Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		todo := pb.Todo{}
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

	return &pb.TodoResponse{
		Items: todos,
	}, nil
}

// CreateTodo implements pb.TodoServiceServer
func (t *todoService) CreateTodo(ctx context.Context, todo *pb.Todo) (*pb.CreateTodoResponse, error) {
	_, err := t.db.Collection("todos").InsertOne(context.TODO(), todo)

	if err != nil {
		return nil, err
	}

	return &pb.CreateTodoResponse{
		Message: "Todo Created :::",
	}, nil
}
