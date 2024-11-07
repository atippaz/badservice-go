package controller

import "github.com/labstack/echo/v4"

type MatchController interface {
	Insert(ctx echo.Context) error
	FindOne(ctx echo.Context) error
	FindAll(ctx echo.Context) error
}
