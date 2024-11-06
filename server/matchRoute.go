package server

import (
	_adminController "bad-service-go/pkg/room/controller"
	_adminRepository "bad-service-go/pkg/room/repository"
	_adminService "bad-service-go/pkg/room/service"
)

func (s *echoServer) InitMatchRouter() {

	adminRepository := _adminRepository.NewRoomRepositoryImpl(s.app.Logger, s.db)
	adminService := _adminService.NewAdminServiceImpl(adminRepository)
	adminController := _adminController.RoomController(adminService)

	s.app.POST("/team", adminController.Insert)
	s.app.GET("/team", adminController.GetAll)
	s.app.GET("/team/:teamId", adminController.GetByID)
	s.app.GET("/deleteTeam/:teamId", adminController.Update)
}
