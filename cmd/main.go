package main

import (
	"log"
	"net"

	"github.com/satioO/todo-grpc/api"
	"github.com/satioO/todo-grpc/service"
	"google.golang.org/grpc"
)

func main() {
	ls, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatal("Failed to start server")
	}

	server := grpc.NewServer()
	todo := service.NewTodoService()
	api.RegisterTodoServiceServer(server, todo)

	if err := server.Serve(ls); err != nil {
		log.Fatal("Failed to start server")
	}
}
