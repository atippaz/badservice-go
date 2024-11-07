package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	RoomName        string    `bson:"roomName"`        // Room name as a string
	RoomCreateOn    time.Time `bson:"roomCreateOn"`    // Creation date as a Time object
	RoomDescription string    `bson:"roomDescription"` // Room description as a string
}
