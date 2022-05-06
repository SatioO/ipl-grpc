package players_svc

import (
	"context"
	"log"

	player_pb "github.com/satioO/todo-grpc/pkg/players/pb"
	teams_repo "github.com/satioO/todo-grpc/pkg/teams/repository"
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

	log.Println(teams)

	return nil, nil
}
