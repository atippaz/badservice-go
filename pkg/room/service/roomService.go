package service

import (
	_AdminModel "bad-service-go/pkg/room/model"
)

type RoomService interface {
	Insert(itemFilter *_AdminModel.RoomInsertRequest) (*_AdminModel.RoomResult, error)
	FindById(AdminId string) (*_AdminModel.RoomResult, error)
	FindAll() (*[]_AdminModel.RoomResult, error)
	GetRoomId()(*[]string,error)
	DeleteRoomById(roomId string)(error)
}
