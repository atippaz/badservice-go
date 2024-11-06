package server

import (
	_adminController "github.com/atippaz/isekai-shop/pkg/admin/controller"
	_adminRepository "github.com/atippaz/isekai-shop/pkg/admin/repository"
	_adminService "github.com/atippaz/isekai-shop/pkg/admin/service"
)

func (s *echoServer) InittRoomRouter() {

	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.app.Logger, s.db)
	adminService := _adminService.NewAdminServiceImpl(adminRepository)
	adminController := _adminController.NewAdminControllerImpl(adminService)

	s.app.POST("/room", adminController.Insert)
	s.app.GET("/room", adminController.GetAll)
	s.app.GET("/room/:roomId", adminController.GetAll)
	s.app.GET("/getRoomId", adminController.GetByID)
	s.app.GET("/deleteRoom/:roomId", adminController.Update)
}
