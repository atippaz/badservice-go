package service

import (
	"bad-service-go/entities"
	_roomModel "bad-service-go/pkg/room/model"
	_roomRepository "bad-service-go/pkg/room/repository"
	"time"
)

type roomServiceImpl struct {
	roomRepository _roomRepository.RoomRepository
}

func NewroomServiceImpl(roomRepository _roomRepository.RoomRepository) RoomService {
	return &roomServiceImpl{
		roomRepository: roomRepository,
	}
}
func (r *roomServiceImpl) DeleteRoomById(roomId string)(error){
	err:=	r.roomRepository.DeleteById(roomId)	
	if err!= nil {
		return err
	}
	return nil
}
func (r *roomServiceImpl) GetRoomId()(*[]string,error){
	value,error := r.roomRepository.FindAll()
	var selectedUsers []string
    for _, user := range value {
        selectedUsers = append(selectedUsers, user.ID.Hex())
    }
	if error !=nil {
		return nil ,error
	}
	return &selectedUsers, nil
}

func (r *roomServiceImpl) Insert(roomRequest *_roomModel.RoomInsertRequest) (*_roomModel.RoomResult, error) {
	const layout = "2006-W02"
	parsedTime, err :=time.Parse(layout, roomRequest.RoomCreateOn)
	if err != nil {
		return nil, err
	}
	newroom := entities.Room{
		RoomName: roomRequest.RoomName,
		RoomCreateOn:parsedTime ,
		RoomDescription: roomRequest.RoomDescription,
	}
	data, erro := r.roomRepository.Insert(newroom)
	if erro != nil {
		return nil, erro
	}
	return &_roomModel.RoomResult{
		RoomId: data.ID.Hex(),
		RoomName: data.RoomName,
		RoomDescription: data.RoomDescription,
		RoomCreateOn: data.RoomCreateOn,
		}, nil
}

func (r *roomServiceImpl) FindAll() (*[]_roomModel.RoomResult, error) {
	result, err := r.roomRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var res []_roomModel.RoomResult 
	for _,data := range result{
		res = append(res, _roomModel.RoomResult{
		RoomId: data.ID.Hex(),
		RoomName: data.RoomName,
		RoomDescription: data.RoomDescription,
		RoomCreateOn: data.RoomCreateOn,
		})
	}
	return &res,nil
}
func (r *roomServiceImpl) FindById(roomId string) (*_roomModel.RoomResult, error) {
	data, error := r.roomRepository.FindById(roomId)
	if error != nil {
		return nil, error
	}
	return &_roomModel.RoomResult{
		RoomId: data.ID.Hex(),
		RoomName: data.RoomName,
		RoomDescription: data.RoomDescription,
		RoomCreateOn: data.RoomCreateOn,
		}, nil
}
