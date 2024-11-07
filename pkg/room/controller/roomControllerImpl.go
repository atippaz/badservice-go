package controller

import (

	// validation "bad-service-go/pkg/room"
	// _roomModel "bad-service-go/pkg/room/model"
	_roomService "bad-service-go/pkg/room/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type roomControllerImpl struct {
	roomService _roomService.RoomService
}

func NewroomControllerImpl(roomService _roomService.RoomService) RoomController {
	return &roomControllerImpl{
		roomService: roomService,
	}
}

func (s *roomControllerImpl) DeleteRoomById(ctx echo.Context) error {
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

func (s *roomControllerImpl) FindAll(ctx echo.Context) error {
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
func (s *roomControllerImpl) FindOne(ctx echo.Context) error {
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

func (s *roomControllerImpl) GetRoomId(ctx echo.Context) error {
	value,error := s.roomService.GetRoomId()


	if error !=nil {
				// return custom.CustomError(ctx, http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, value)
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
	
}
