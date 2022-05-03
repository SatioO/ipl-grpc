package teams_svc

import (
	context "context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/satioO/todo-grpc/pkg/teams"
	"github.com/satioO/todo-grpc/pkg/teams/model"
	teams_pb "github.com/satioO/todo-grpc/pkg/teams/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type teamService struct {
	teams_pb.UnimplementedTeamServiceServer
	db         *mongo.Database
	collection string
}

func NewTeamService(db *mongo.Database, collection string) teams_pb.TeamServiceServer {
	return &teamService{
		db:         db,
		collection: collection,
	}
}

// CreateTeams implements teams_pb.TeamServiceServer
func (t *teamService) CreateTeams(ctx context.Context, rq *teams_pb.CreateTeamsRequest) (*teams_pb.CreateTeamsResponse, error) {
	ch := make(chan *model.Team)

	res, err := http.Get(teams.GET_TEAM_URL)

	if err != nil {
		return nil, err
	}

	var teams model.TeamsResponse
	err = json.NewDecoder(res.Body).Decode(&teams)
	if err != nil {
		return nil, err
	}

	for _, team := range teams.Items {
		go func(url string) {
			team, err := t.FetchTeamDetails(url)
			if err != nil {
				log.Println("error getting teams", err)
			}

			ch <- team
		}(team.URL)
	}

	var output []interface{}
	for i := 0; i < teams.Count; i++ {
		comm := <-ch
		output = append(output, comm)
	}

	_, err = t.db.Collection(t.collection).InsertMany(ctx, output)
	if err != nil {
		return nil, err
	}

	return &teams_pb.CreateTeamsResponse{
		Message: "Inserted: 1",
	}, nil
}

// GetTeams implements teams_pb.TeamServiceServer
func (t *teamService) GetTeams(ctx context.Context, rq *teams_pb.GetTeamRequest) (*teams_pb.GetTeamResponse, error) {
	teams := []*teams_pb.Team{}
	cursor, err := t.db.Collection(t.collection).Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		team := model.Team{}
		cursor.Decode(&team)

		teams = append(teams, &teams_pb.Team{
			Id:    team.ID,
			Name:  team.Name,
			Color: team.Color,
			Abbr:  team.Abbr,
			Logo:  team.Logos[0].Href,
		})
	}

	return &teams_pb.GetTeamResponse{
		Items: teams,
	}, nil
}

func (t *teamService) FetchTeamDetails(url string) (*model.Team, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	var team *model.Team
	err = json.NewDecoder(res.Body).Decode(&team)
	if err != nil {
		return nil, err
	}

	return team, nil
}
