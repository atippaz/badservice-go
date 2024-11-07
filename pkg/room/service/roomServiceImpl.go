package service

import (
	"bad-service-go/entities"
	_roomModel "bad-service-go/pkg/room/model"
	_roomRepository "bad-service-go/pkg/room/repository"
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
	newroom := entities.Room{
		// Email:  roomRequest.Email,
		// Avatar: roomRequest.Avatar,
		// Name:   roomRequest.Name,
	}
	_, err := r.roomRepository.Insert(newroom)
	if err != nil {
		return nil, err
	}
	return &_roomModel.RoomResult{
		// ID:     result.ID,
		// Email:  result.Email,
		// Name:   result.Name,
		// Avatar: result.Avatar,
	}, nil
}
func (r *roomServiceImpl) FindAll() (*[]_roomModel.RoomResult, error) {
	_, error := r.roomRepository.FindAll()
	if error != nil {
		return nil, error
	}
	// return &_roomModel.RoomResult{
	// 	// ID:     result.ID,
	// 	// Email:  result.Email,
	// 	// Name:   result.Name,
	// 	// Avatar: result.Avatar,
	// }, nil
	return nil,nil
}
func (r *roomServiceImpl) FindById(roomId string) (*_roomModel.RoomResult, error) {
	_, error := r.roomRepository.FindById(roomId)
	if error != nil {
		return nil, error
	}
	return &_roomModel.RoomResult{
		// ID:     result.ID,
		// Email:  result.Email,
		// Name:   result.Name,
		// Avatar: result.Avatar,
	}, nil
}
