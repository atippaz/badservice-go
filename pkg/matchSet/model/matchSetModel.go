package model

import "bad-service-go/entities"

type (
	MatchSet struct {
		ID     string `json:"id"`
		Email  string `json:"email"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}

	MatchSetInsertRequest struct {
		Email  string `json:"email"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}

	MatchSetResult struct {
		RoomId      string `json:"roomId"`
		SetId       string `json:"setId"`
		TeamName    string `json:"teamName"`
		CourtNumber int    `json:"courtNumber"`
		AllTeam     []entities.TeamSet `json:"allTeam"`
		WinScore    int    `json:"winScore"`
		TeamLimit   int    `json:"teamLimit"`
		WinStreak   int    `json:"winStreak"`
	}
)
