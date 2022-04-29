package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/satioO/todo-grpc/pkg/todo/pb"
	"github.com/satioO/todo-grpc/pkg/todo/service"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ls, err := net.Listen("tcp", fmt.Sprintf(":%v", os.Getenv("SERVICE_PORT")))

	if err != nil {
		log.Fatal("Failed to start server")
	}

	conn, err := GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	todo := service.NewTodoService(conn.Database(os.Getenv("DB_NAME")))
	pb.RegisterTodoServiceServer(server, todo)

	if err := server.Serve(ls); err != nil {
		log.Fatal("Failed to start server")
	}
}
