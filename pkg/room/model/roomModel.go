package model

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
		ID     string `json:"id"`
		Email  string `json:"email"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}
)
