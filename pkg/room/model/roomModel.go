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
		Email  string `json:"email"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}

	RoomResult struct {
		RoomId          string    `json:"roomId"`
		RoomName        string    `json:"roomName"`
		RoomDescription string    `json:"roomDescription"`
		RoomCreateOn    time.Time `json:"roomCreateOn"`
	}
)
