package model

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
		ID     string `json:"id"`
		Email  string `json:"email"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}
)
