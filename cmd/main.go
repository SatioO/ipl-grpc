package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/satioO/todo-grpc/api"
	"github.com/satioO/todo-grpc/service"
	"google.golang.org/grpc"
)

func main() {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
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

	server := grpc.NewServer(withUnaryServerInterceptor())
	todo := service.NewTodoService(conn.Database("todo-db"))
	api.RegisterTodoServiceServer(server, todo)

	if err := server.Serve(ls); err != nil {
		log.Fatal("Failed to start server")
	}
}
