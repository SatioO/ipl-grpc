package teams_repo

import (
	"context"

	"github.com/satioO/todo-grpc/pkg/teams/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TeamsRepo struct {
	collection *mongo.Collection
}

func NewTeamsRepo(collection *mongo.Collection) *TeamsRepo {
	return &TeamsRepo{
		collection: collection,
	}
}

func (t *TeamsRepo) GetTeams() ([]model.Team, error) {
	var teams []model.Team
	cursor, err := t.collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		team := model.Team{}
		cursor.Decode(&team)
		teams = append(teams, team)
	}

	return teams, nil
}

func (t *TeamsRepo) InsertTeams(teams []interface{}) (*mongo.InsertManyResult, error) {
	return t.collection.InsertMany(context.TODO(), teams)
}
