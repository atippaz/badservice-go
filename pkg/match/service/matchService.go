package service

import (
	_matchModel "bad-service-go/pkg/match/model"
)

type MatchService interface {
	Insert(itemFilter *_matchModel.MatchInsertRequest) (*_matchModel.MatchResult, error)
	FindById(matchId string) (*_matchModel.MatchResult, error)
	FindAll() (*[]_matchModel.MatchResult, error)
}
