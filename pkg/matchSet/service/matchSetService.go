package service

import (
	_AdminModel "bad-service-go/pkg/matchSet/model"
)

type MatchSetService interface {
	Insert(itemFilter *_AdminModel.MatchSetInsertRequest) (*_AdminModel.MatchSetResult, error)
	FindById(AdminId string) (*_AdminModel.MatchSetResult, error)
	FindAll() (*[]_AdminModel.MatchSetResult, error)
	DeleteById(matchSetId string) ( error)

}
