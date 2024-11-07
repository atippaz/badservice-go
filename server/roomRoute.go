package server

import (
	_roomController "bad-service-go/pkg/room/controller"
	_roomRepository "bad-service-go/pkg/room/repository"
	_roomService "bad-service-go/pkg/room/service"
)

func (s *echoServer) InittRoomRouter() {

	roomRepository := _roomRepository.NewRoomRepositoryImpl(s.app.Logger, s.db)
	roomService := _roomService.NewroomServiceImpl(roomRepository)
	roomController := _roomController.NewroomControllerImpl(roomService)

	s.app.POST("/room", roomController.Insert)
	s.app.GET("/room", roomController.FindAll)
	s.app.GET("/room/:roomId", roomController.FindOne)
	s.app.GET("/getRoomId", roomController.GetRoomId)
	s.app.GET("/deleteRoom/:roomId", roomController.DeleteRoomById)
}
