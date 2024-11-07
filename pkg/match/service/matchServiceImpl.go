package service

import (
	"bad-service-go/entities"
	_matchModel "bad-service-go/pkg/match/model"
	_matchRepository "bad-service-go/pkg/match/repository"
)

type matchServiceImpl struct {
	matchRepository _matchRepository.MatchRepository
}

func NewmatchServiceImpl(matchRepository _matchRepository.MatchRepository) MatchService {
	return &matchServiceImpl{
		matchRepository: matchRepository,
	}
}

func (r *matchServiceImpl) DeleteById(matchId string) ( error) {
	err := r.matchRepository.DeleteById(matchId)
	if err != nil {
		return err
	}
	return nil
}

func (r *matchServiceImpl) Insert(matchRequest *_matchModel.MatchInsertRequest) (*_matchModel.MatchResult, error) {
	newmatch := entities.Match{
		// Email:  matchRequest.Email,
		// Avatar: matchRequest.Avatar,
		// Name:   matchRequest.Name,
	}
	_, err := r.matchRepository.Insert(newmatch)
	if err != nil {
		return nil, err
	}
	return &_matchModel.MatchResult{
		// ID:     result.ID,
		// Email:  result.Email,
		// Name:   result.Name,
		// Avatar: result.Avatar,
	}, nil
}

func (r *matchServiceImpl) FindAll() (*[]_matchModel.MatchResult, error) {
	result, error := r.matchRepository.FindAll()
	if error != nil {
		return nil, error
	}
	var res []_matchModel.MatchResult
	for _,data  := range result{
		res = append(res, _matchModel.MatchResult{
		RoomId: data.RoomID,
		TeamId: data.ID.Hex(),
		TeamName: data.SetName,
		CourtNumber:data.CourtNumber,
		AllTeam: data.AllTeam,
		WinScore: data.WinScore,
		TeamLimit: data.TeamLimit,
		WinStreak: data.WinStreak,
		})
	}
	return &res, nil
}

func (r *matchServiceImpl) FindById(matchId string) (*_matchModel.MatchResult, error) {
	data, error := r.matchRepository.FindById(matchId)
	if error != nil {
		return nil, error
	}
	return &_matchModel.MatchResult{
		RoomId: data.RoomID,
		TeamId: data.ID.Hex(),
		TeamName: data.SetName,
		CourtNumber:data.CourtNumber,
		AllTeam: data.AllTeam,
		WinScore: data.WinScore,
		TeamLimit: data.TeamLimit,
		WinStreak: data.WinStreak,
		}, nil
}
