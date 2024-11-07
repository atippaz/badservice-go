package repository

import (
	"bad-service-go/databases"
	"bad-service-go/entities"
	"context"
	"errors"

	// _matchSetException "bad-service-go/pkg/matchSet/exception"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatchSetRepositoryImpl struct {
	logger     echo.Logger
	collection *mongo.Collection
}

func NewMatchSetRepositoryImpl(
	logger echo.Logger,
	db databases.Database,
) MatchSetRepository {
	return &MatchSetRepositoryImpl{
		logger:     logger,
		collection: db.Connect().Collection("matchSets"),
	}
}

func (r *MatchSetRepositoryImpl) Insert(matchSetEntity entities.MatchSet) (*entities.MatchSet, error) {
	_, err := r.collection.InsertOne(context.TODO(), matchSetEntity)
	if err != nil {
		return nil, err
	}
	return &matchSetEntity, nil
}

func (r *MatchSetRepositoryImpl) DeleteById(matchSetId string) ( error) {
	filter := bson.M{"id": matchSetId}
	_, err := r.collection.DeleteOne(context.TODO(), filter)
	return err
}

func (r *MatchSetRepositoryImpl) FindById(matchSetId string) (*entities.MatchSet, error) {
	filter := bson.M{"matchSetId": matchSetId}

	var matchSet entities.MatchSet
	err := r.collection.FindOne(context.TODO(), filter).Decode(&matchSet)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("matchSet not found")
		}
		return nil, err
	}
	return &matchSet, nil
}
func (r *MatchSetRepositoryImpl) FindAll() ([]entities.MatchSet, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var matchSets []entities.MatchSet
	for cursor.Next(context.TODO()) {
		var matchSet entities.MatchSet
		if err := cursor.Decode(&matchSet); err != nil {
			return nil, err
		}
		matchSets = append(matchSets, matchSet)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return matchSets, nil
}
