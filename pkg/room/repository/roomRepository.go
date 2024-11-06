package repository

import (
	"bad-service-go/entities"
)

type RoomRepository interface {
	Insert(entities.Room) (*entities.Room, error)
	FindById(adminId string) (*entities.Room, error)
	FindAll() ([]entities.Room, error)
}
