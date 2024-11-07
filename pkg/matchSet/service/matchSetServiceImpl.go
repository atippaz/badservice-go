package service

import (
	"bad-service-go/entities"
	_matchSetModel "bad-service-go/pkg/matchSet/model"
	_matchSetRepository "bad-service-go/pkg/matchSet/repository"
)

type matchSetServiceImpl struct {
	matchSetRepository _matchSetRepository.MatchSetRepository
}

func NewmatchSetServiceImpl(matchSetRepository _matchSetRepository.MatchSetRepository) MatchSetService {
	return &matchSetServiceImpl{
		matchSetRepository: matchSetRepository,
	}
}
 
func (r *matchSetServiceImpl) DeleteById(matchSetId string) error {
	err := r.matchSetRepository.DeleteById(matchSetId)
	if err != nil {
		return err
	}
	return nil
}
func (r *matchSetServiceImpl) Insert(matchSetRequest *_matchSetModel.MatchSetInsertRequest) (*_matchSetModel.MatchSetResult, error) {
	newmatchSet := entities.MatchSet{
		// Email:  matchSetRequest.Email,
		// Avatar: matchSetRequest.Avatar,
		// Name:   matchSetRequest.Name,
	}
	_, err := r.matchSetRepository.Insert(newmatchSet)
	if err != nil {
		return nil, err
	}
	return &_matchSetModel.MatchSetResult{
		// ID:     result.ID,
		// Email:  result.Email,
		// Name:   result.Name,
		// Avatar: result.Avatar,
	}, nil
}
func (r *matchSetServiceImpl) FindAll() (*[]_matchSetModel.MatchSetResult, error) {
	result, error := r.matchSetRepository.FindAll()
	if error != nil {
		return nil, error
	}
	var res []_matchSetModel.MatchSetResult
	for _,data := range result{
		res = append(res, _matchSetModel.MatchSetResult{
			RoomId: data.RoomID,
			SetId: data.ID.Hex(),
			TeamName: data.SetName,
			CourtNumber: data.CourtNumber,
			AllTeam: data.AllTeam,
			WinScore: data.WinScore,
			TeamLimit: data.TeamLimit,
			WinStreak: data.WinStreak,
		})
	}
	return &res,nil
}
func (r *matchSetServiceImpl) FindById(matchSetId string) (*_matchSetModel.MatchSetResult, error) {
	data, error := r.matchSetRepository.FindById(matchSetId)
	if error != nil {
		return nil, error
	}
	return &_matchSetModel.MatchSetResult{
		RoomId: data.RoomID,
		SetId: data.ID.Hex(),
		TeamName: data.SetName,
		CourtNumber: data.CourtNumber,
		AllTeam: data.AllTeam,
		WinScore: data.WinScore,
		TeamLimit: data.TeamLimit,
		WinStreak: data.WinStreak,
	}, nil
}
