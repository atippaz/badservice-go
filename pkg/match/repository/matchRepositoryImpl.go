package repository

import (
	"bad-service-go/databases"
	"bad-service-go/entities"
	"context"
	"errors"

	// _matchException "bad-service-go/pkg/match/exception"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatchRepositoryImpl struct {
	logger     echo.Logger
	collection *mongo.Collection
}

func NewmatchRepositoryImpl(
	logger echo.Logger,
	db databases.Database,
) MatchRepository {
	return &MatchRepositoryImpl{
		logger:     logger,
		collection: db.Connect().Collection("matchs"),
	}
}
func (r *MatchRepositoryImpl) Insert(matchEntity entities.Match) (*entities.Match, error) {
	_, err := r.collection.InsertOne(context.TODO(), matchEntity)
	if err != nil {
		return nil, err
	}
	return &matchEntity, nil
}

func (r *MatchRepositoryImpl) DeleteById(matchId string) ( error) {
	filter := bson.M{"id": matchId}
	_, err := r.collection.DeleteOne(context.TODO(), filter)
	return err
}
func (r *MatchRepositoryImpl) FindById(matchId string) (*entities.Match, error) {
	filter := bson.M{"matchId": matchId}

	var match entities.Match
	err := r.collection.FindOne(context.TODO(), filter).Decode(&match)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("match not found")
		}
		return nil, err
	}
	return &match, nil
}
func (r *MatchRepositoryImpl) FindAll() ([]entities.Match, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var matchs []entities.Match
	for cursor.Next(context.TODO()) {
		var match entities.Match
		if err := cursor.Decode(&match); err != nil {
			return nil, err
		}
		matchs = append(matchs, match)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return matchs, nil
}
