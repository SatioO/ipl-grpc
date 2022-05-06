package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/satioO/todo-grpc/pkg/players"
	player_pb "github.com/satioO/todo-grpc/pkg/players/pb"
	shared_model "github.com/satioO/todo-grpc/pkg/shared/model"
	teams_repo "github.com/satioO/todo-grpc/pkg/teams/repo"
)

type playerService struct {
	player_pb.UnimplementedPlayerServiceServer
	teams_repo *teams_repo.TeamsRepo
}

func NewPlayerService(teams_repo *teams_repo.TeamsRepo) player_pb.PlayerServiceServer {
	return &playerService{
		teams_repo: teams_repo,
	}
}

// GetPlayers implements pb.PlayerServiceServer
func (*playerService) GetPlayers(ctx context.Context, rq *player_pb.TeamPlayerRequest) (*player_pb.TeamPlayerResponse, error) {
	first := []*player_pb.Player{}
	second := []*player_pb.Player{}

	return &player_pb.TeamPlayerResponse{
		First:  first,
		Second: second,
	}, nil
}

func map_player_type(record []string) player_pb.PlayerType {
	switch record[1] {
	default:
		return player_pb.PlayerType(player_pb.PlayerType_AllRounder)
	}
}

// UploadPlayers implements pb.PlayerServiceServer
func (p *playerService) CreatePlayers(ctx context.Context, rq *player_pb.UploadPlayersRequest) (*player_pb.UploadPlayersResponse, error) {
	teams, err := p.teams_repo.GetTeams()
	if err != nil {
		return nil, err
	}

	for _, team := range teams {
		go func(teamId string) {
			res, err := GetAthletes(teamId)
			if err != nil {
				log.Println("error getting teams", err)
			}

			log.Println(res)
		}(team.ID)
	}

	return &player_pb.UploadPlayersResponse{}, nil
}

func GetAthletes(teamId string) (*shared_model.Response, error) {
	res, err := http.Get(players.GET_ATHLETES_URL(os.Getenv("SEASON_ID"), teamId))

	if err != nil {
		return nil, err
	}

	var response shared_model.Response
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
