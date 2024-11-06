package server

import (
	_adminController "github.com/atippaz/isekai-shop/pkg/admin/controller"
	_adminRepository "github.com/atippaz/isekai-shop/pkg/admin/repository"
	_adminService "github.com/atippaz/isekai-shop/pkg/admin/service"
)

func (s *echoServer) InitMatchSetRouter() {

	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.app.Logger, s.db)
	adminService := _adminService.NewAdminServiceImpl(adminRepository)
	adminController := _adminController.NewAdminControllerImpl(adminService)

	s.app.GET("/set", adminController.Insert)
	s.app.GET("/set/:setId", adminController.GetAll)
	s.app.GET("/deleteSet/:setId", adminController.Update)
}
