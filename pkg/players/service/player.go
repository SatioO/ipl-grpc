package players

import (
	context "context"

	"github.com/satioO/todo-grpc/pkg/players/pb"
)

type playerService struct {
	pb.UnimplementedPlayerServiceServer
}

// GetPlayers implements pb.PlayerServiceServer
func (*playerService) GetPlayers(context.Context, *pb.TeamPlayerRequest) (*pb.TeamPlayerResponse, error) {
	panic("unimplemented")
}

func NewPlayerService() pb.PlayerServiceServer {
	return &playerService{}
}
