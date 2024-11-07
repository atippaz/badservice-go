package controller

import "github.com/labstack/echo/v4"

type RoomController interface {
	Insert(ctx echo.Context) error
	FindOne(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	GetRoomId(ctx echo.Context) error
	DeleteRoomById(ctx echo.Context) error

}
