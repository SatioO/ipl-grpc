package teams_svc

import (
	context "context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/satioO/todo-grpc/pkg/teams"
	"github.com/satioO/todo-grpc/pkg/teams/model"
	teams_pb "github.com/satioO/todo-grpc/pkg/teams/pb"
	teams_repo "github.com/satioO/todo-grpc/pkg/teams/repository"
)

type TeamService struct {
	teams_pb.UnimplementedTeamServiceServer
	repo *teams_repo.TeamsRepo
}

func NewTeamService(repo *teams_repo.TeamsRepo) teams_pb.TeamServiceServer {
	return &TeamService{
		repo: repo,
	}
}

// CreateTeams implements teams_pb.TeamServiceServer
func (t *TeamService) CreateTeams(ctx context.Context, rq *teams_pb.CreateTeamsRequest) (*teams_pb.CreateTeamsResponse, error) {
	ch := make(chan *model.Team)
	defer close(ch)

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

	_, err = t.repo.InsertTeams(output)

	if err != nil {
		return nil, err
	}

	return &teams_pb.CreateTeamsResponse{
		Message: "Inserted: 1",
	}, nil
}

// GetTeams implements teams_pb.TeamServiceServer
func (t *TeamService) GetTeams(ctx context.Context, rq *teams_pb.GetTeamRequest) (*teams_pb.GetTeamResponse, error) {
	teams := []*teams_pb.Team{}
	entity, err := t.repo.GetTeams()

	if err != nil {
		return nil, err
	}

	for _, team := range entity {
		teams = append(teams, &teams_pb.Team{
			Id:    team.ID,
			Name:  team.Name,
			Abbr:  team.Abbr,
			Color: team.Color,
			Logo:  team.Logos[0].Href,
		})
	}

	return &teams_pb.GetTeamResponse{
		Items: teams,
	}, nil
}

func (t *TeamService) FetchTeamDetails(url string) (*model.Team, error) {
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
