package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	player_pb "github.com/satioO/todo-grpc/pkg/players/pb"
	players_svc "github.com/satioO/todo-grpc/pkg/players/service"
	teams_pb "github.com/satioO/todo-grpc/pkg/teams/pb"
	teams_svc "github.com/satioO/todo-grpc/pkg/teams/service"
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

	player := players_svc.NewPlayerService(conn.Database(os.Getenv("DB_NAME")), "players")
	player_pb.RegisterPlayerServiceServer(server, player)

	team := teams_svc.NewTeamService(conn.Database(os.Getenv("DB_NAME")), "teams")
	teams_pb.RegisterTeamServiceServer(server, team)

	if err := server.Serve(ls); err != nil {
		log.Fatal("Failed to start server")
	}
}
