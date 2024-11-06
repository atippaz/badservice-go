package controller

import "github.com/labstack/echo/v4"

type RoomController interface {
	Insert(ctx echo.Context) error
}
