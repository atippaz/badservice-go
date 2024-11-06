package entities

import "time"

type Room struct {
	RoomName        string    `bson:"roomName"`        // Room name as a string
	RoomCreateOn    time.Time `bson:"roomCreateOn"`    // Creation date as a Time object
	RoomDescription string    `bson:"roomDescription"` // Room description as a string
}
