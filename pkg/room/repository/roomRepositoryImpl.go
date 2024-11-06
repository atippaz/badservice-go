package repository

import (
	"bad-service-go/databases"
	"bad-service-go/entities"
	"context"
	"errors"

	// _roomException "bad-service-go/pkg/room/exception"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomRepositoryImpl struct {
	logger     echo.Logger
	collection *mongo.Collection
}

func NewRoomRepositoryImpl(
	logger echo.Logger,
	db databases.Database,
) RoomRepository {
	return &RoomRepositoryImpl{
		logger:     logger,
		collection: db.Connect().Collection("rooms"),
	}
}
func (r *RoomRepositoryImpl) Insert(roomEntity entities.Room) (*entities.Room, error) {
	_, err := r.collection.InsertOne(context.TODO(), roomEntity)
	if err != nil {
		return nil, err
	}
	return &roomEntity, nil
}

func (r *RoomRepositoryImpl) FindById(roomId string) (*entities.Room, error) {
	filter := bson.M{"roomId": roomId}

	var room entities.Room
	err := r.collection.FindOne(context.TODO(), filter).Decode(&room)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("room not found")
		}
		return nil, err
	}
	return &room, nil
}
func (r *RoomRepositoryImpl) FindAll() ([]entities.Room, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var rooms []entities.Room
	for cursor.Next(context.TODO()) {
		var room entities.Room
		if err := cursor.Decode(&room); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}
