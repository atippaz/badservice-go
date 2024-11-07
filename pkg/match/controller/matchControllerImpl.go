package controller

import (

	// validation "bad-service-go/pkg/match"
	// _matchModel "bad-service-go/pkg/match/model"
	_matchService "bad-service-go/pkg/match/service"

	"github.com/labstack/echo/v4"
)

type matchControllerImpl struct {
	matchService _matchService.MatchService
}

func NewMatchControllerImpl(matchService _matchService.MatchService) MatchController {
	return &matchControllerImpl{
		matchService: matchService,
	}
}

func (s *matchControllerImpl) DeleteById(ctx echo.Context) error {
	// itemCreateRequest := new(_matchModel.matchInsertRequest)
	// customEchoRequest := validation.NewCustomEchoRequest(ctx)
	// if err := customEchoRequest.Bind(itemCreateRequest); err != nil {
	// 	return custom.CustomError(ctx, http.StatusInternalServerError, err)
	// }

	// newItem, err := s.matchService.Insert(itemCreateRequest)
	// if err != nil {
	// 	return custom.CustomError(ctx, http.StatusInternalServerError, err)
	// }
	// return ctx.JSON(http.StatusOK, newItem)
	return nil
}

func (s *matchControllerImpl) Insert(ctx echo.Context) error {
	// itemCreateRequest := new(_matchModel.matchInsertRequest)
	// customEchoRequest := validation.NewCustomEchoRequest(ctx)
	// if err := customEchoRequest.Bind(itemCreateRequest); err != nil {
	// 	return custom.CustomError(ctx, http.StatusInternalServerError, err)
	// }

	// newItem, err := s.matchService.Insert(itemCreateRequest)
	// if err != nil {
	// 	return custom.CustomError(ctx, http.StatusInternalServerError, err)
	// }
	// return ctx.JSON(http.StatusOK, newItem)
	return nil
}

func (s *matchControllerImpl) FindAll(ctx echo.Context) error {
	// itemCreateRequest := new(_matchModel.matchInsertRequest)
	// customEchoRequest := validation.NewCustomEchoRequest(ctx)
	// if err := customEchoRequest.Bind(itemCreateRequest); err != nil {
	// 	return custom.CustomError(ctx, http.StatusInternalServerError, err)
	// }

	// newItem, err := s.matchService.Insert(itemCreateRequest)
	// if err != nil {
	// 	return custom.CustomError(ctx, http.StatusInternalServerError, err)
	// }
	// return ctx.JSON(http.StatusOK, newItem)
	return nil
}

func (s *matchControllerImpl) FindOne(ctx echo.Context) error {
	// itemCreateRequest := new(_matchModel.matchInsertRequest)
	// customEchoRequest := validation.NewCustomEchoRequest(ctx)
	// if err := customEchoRequest.Bind(itemCreateRequest); err != nil {
	// 	return custom.CustomError(ctx, http.StatusInternalServerError, err)
	// }

	// newItem, err := s.matchService.Insert(itemCreateRequest)
	// if err != nil {
	// 	return custom.CustomError(ctx, http.StatusInternalServerError, err)
	// }
	// return ctx.JSON(http.StatusOK, newItem)
	return nil
}
