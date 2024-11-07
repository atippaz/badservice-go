package server

import (
	_matchSetController "bad-service-go/pkg/matchSet/controller"
	_matchSetRepository "bad-service-go/pkg/matchSet/repository"
	_matchSetService "bad-service-go/pkg/matchSet/service"
)

func (s *echoServer) InitMatchSetRouter() {

	matchSetRepository := _matchSetRepository.NewMatchSetRepositoryImpl(s.app.Logger, s.db)
	matchSetService := _matchSetService.NewmatchSetServiceImpl(matchSetRepository)
	matchSetController := _matchSetController.NewmatchSetControllerImpl(matchSetService)

	s.app.GET("/set", matchSetController.Insert)
	s.app.GET("/set/:setId", matchSetController.FindAll)
	s.app.GET("/deleteSet/:setId", matchSetController.FindAll)
}
