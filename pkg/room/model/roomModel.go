package model

type (
	Admin struct {
		ID     string `json:"id"`
		Email  string `json:"email"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}

	AdminInsertRequest struct {
		Email  string `json:"email"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}

	AdminResult struct {
		ID     string `json:"id"`
		Email  string `json:"email"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}
)
