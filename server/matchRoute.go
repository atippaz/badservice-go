package server

import (
	_matchController "bad-service-go/pkg/match/controller"
	_matchRepository "bad-service-go/pkg/match/repository"
	_matchService "bad-service-go/pkg/match/service"
)

func (s *echoServer) InitMatchRouter() {

	matchRepository := _matchRepository.NewmatchRepositoryImpl(s.app.Logger, s.db)
	matchService := _matchService.NewmatchServiceImpl(matchRepository)
	matchController := _matchController.NewMatchControllerImpl(matchService)

	s.app.POST("/team", matchController.Insert)
	s.app.GET("/team", matchController.FindAll)
	s.app.GET("/team/:teamId", matchController.FindAll)
	s.app.GET("/deleteTeam/:teamId", matchController.DeleteById)
}
