package repository

import (
	"bad-service-go/entities"
)

type MatchRepository interface {
	Insert(entities.Match) (*entities.Match, error)
	FindById(adminId string) (*entities.Match, error)
	FindAll() ([]entities.Match, error)
}
