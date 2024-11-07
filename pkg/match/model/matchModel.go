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

	matchset struct {
		IsSet bool `json:"isSet"`
		LimitSet bool `json:"limitSet"`
	}

	MatchInsertRequest struct {
		Members  *[]string `json:"members"`
		Limit   *int `json:"limit"`
		Start  *int `json:"start"`
		Random *bool `json:"random"`
		MatchSet matchset `json:"matchset"`
		SetName string `json:"setName"`
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
