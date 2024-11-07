package repository

import (
	"bad-service-go/entities"
)

type MatchSetRepository interface {
	Insert(entities.MatchSet) (*entities.MatchSet, error)
	FindById(adminId string) (*entities.MatchSet, error)
	FindAll() ([]entities.MatchSet, error)
}
