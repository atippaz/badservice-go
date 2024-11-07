package model

import "time"

type (
	Room struct {
		ID     string `json:"id"`
		Email  string `json:"email"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}

	RoomInsertRequest struct {
		RoomName string `json:"roomName"` 
		RoomCreateOn string `json:"roomCreateOn"` 
		RoomDescription string `json:"roomDescription"`
	}

	RoomResult struct {
		RoomId          string    `json:"roomId"`
		RoomName        string    `json:"roomName"`
		RoomDescription string    `json:"roomDescription"`
		RoomCreateOn    time.Time `json:"roomCreateOn"`
	}
)
