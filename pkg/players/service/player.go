package players_svc

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	player_pb "github.com/satioO/todo-grpc/pkg/players/pb"
	"go.mongodb.org/mongo-driver/mongo"
)

type playerService struct {
	player_pb.UnimplementedPlayerServiceServer
	db         *mongo.Database
	collection string
}

func NewPlayerService(db *mongo.Database, collection string) player_pb.PlayerServiceServer {
	return &playerService{
		db:         db,
		collection: collection,
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
func (*playerService) UploadPlayers(ctx context.Context, rq *player_pb.UploadPlayersRequest) (*player_pb.UploadPlayersResponse, error) {
	f, err := os.Open("./static/IPL_Data.csv")
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	_, err = csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}
	return nil, nil
}
