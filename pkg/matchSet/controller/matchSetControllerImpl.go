package controller

import (
	_matchSetService "bad-service-go/pkg/matchSet/service"

	"github.com/labstack/echo/v4"
)

type MatchSetControllerImpl struct {
	matchSetService _matchSetService.MatchSetService
}

func NewmatchSetControllerImpl(matchSetService _matchSetService.MatchSetService) MatchSetController {
	return &MatchSetControllerImpl{
		matchSetService: matchSetService,
	}
}

func (s *MatchSetControllerImpl) Insert(ctx echo.Context) error {
	// itemCreateRequest := new(_matchSetModel.matchSetInsertRequest)
	// customEchoRequest := validation.NewCustomEchoRequest(ctx)
	// if err := customEchoRequest.Bind(itemCreateRequest); err != nil {
	// 	return custom.CustomError(ctx, http.StatusInternalServerError, err)
	// }

	// newItem, err := s.matchSetService.Insert(itemCreateRequest)
	// if err != nil {
	// 	return custom.CustomError(ctx, http.StatusInternalServerError, err)
	// }
	// return ctx.JSON(http.StatusOK, newItem)
	return nil
}

func (s *MatchSetControllerImpl) FindAll(ctx echo.Context) error {
	// itemCreateRequest := new(_matchSetModel.matchSetInsertRequest)
	// customEchoRequest := validation.NewCustomEchoRequest(ctx)
	// if err := customEchoRequest.Bind(itemCreateRequest); err != nil {
	// 	return custom.CustomError(ctx, http.StatusInternalServerError, err)
	// }

	// newItem, err := s.matchSetService.Insert(itemCreateRequest)
	// if err != nil {
	// 	return custom.CustomError(ctx, http.StatusInternalServerError, err)
	// }
	// return ctx.JSON(http.StatusOK, newItem)
	return nil
}

func (s *MatchSetControllerImpl) FindOne(ctx echo.Context) error {
	// itemCreateRequest := new(_matchSetModel.matchSetInsertRequest)
	// customEchoRequest := validation.NewCustomEchoRequest(ctx)
	// if err := customEchoRequest.Bind(itemCreateRequest); err != nil {
	// 	return custom.CustomError(ctx, http.StatusInternalServerError, err)
	// }

	// newItem, err := s.matchSetService.Insert(itemCreateRequest)
	// if err != nil {
	// 	return custom.CustomError(ctx, http.StatusInternalServerError, err)
	// }
	// return ctx.JSON(http.StatusOK, newItem)
	return nil
}
