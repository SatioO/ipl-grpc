package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"

	teams_pb "github.com/satioO/todo-grpc/pkg/teams/pb"
	teams_repo "github.com/satioO/todo-grpc/pkg/teams/repo"
	teams_svc "github.com/satioO/todo-grpc/pkg/teams/service"

	player_pb "github.com/satioO/todo-grpc/pkg/players/pb"
	players_svc "github.com/satioO/todo-grpc/pkg/players/service"
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

	// TEAMS
	repo_teams := teams_repo.NewTeamsRepo(conn.Database(os.Getenv("DB_NAME")).Collection("teams"))
	svc_teams := teams_svc.NewTeamService(repo_teams)
	teams_pb.RegisterTeamServiceServer(server, svc_teams)

	// PLAYERS
	svc_players := players_svc.NewPlayerService(repo_teams)
	player_pb.RegisterPlayerServiceServer(server, svc_players)

	if err := server.Serve(ls); err != nil {
		log.Fatal("Failed to start server")
	}
}
