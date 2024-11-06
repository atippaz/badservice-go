package controller

import (

	// validation "bad-service-go/pkg/room"
	// _roomModel "bad-service-go/pkg/room/model"
	_roomService "bad-service-go/pkg/room/service"

	"github.com/labstack/echo/v4"
)

type roomControllerImpl struct {
	roomService _roomService.AdminService
}

func NewroomControllerImpl(roomService _roomService.AdminService) RoomController {
	return &roomControllerImpl{
		roomService: roomService,
	}
}

func (s *roomControllerImpl) Insert(ctx echo.Context) error {
	// itemCreateRequest := new(_roomModel.roomInsertRequest)
	// customEchoRequest := validation.NewCustomEchoRequest(ctx)
	// if err := customEchoRequest.Bind(itemCreateRequest); err != nil {
	// 	return custom.CustomError(ctx, http.StatusInternalServerError, err)
	// }

	// newItem, err := s.roomService.Insert(itemCreateRequest)
	// if err != nil {
	// 	return custom.CustomError(ctx, http.StatusInternalServerError, err)
	// }
	// return ctx.JSON(http.StatusOK, newItem)
	return nil
}
