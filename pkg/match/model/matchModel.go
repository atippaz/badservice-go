package model

import (
	"bad-service-go/entities"
)
type (
	Match struct {
		ID     string `json:"id"`
		Email  string `json:"email"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}

	MatchInsertRequest struct {
		Email  string `json:"email"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}

	MatchResult struct {
		RoomId      string `json:"roomId"`
		TeamId      string `json:"teamId"`
		TeamName    string `json:"teamName"`
		CourtNumber int    `json:"courtNumber"`
		AllTeam     []entities.Team `json:"allTeam"`
		WinScore    int    `json:"winScore"`
		TeamLimit   int    `json:"teamLimit"`
		WinStreak   int    `json:"winStreak"`
	}
)
