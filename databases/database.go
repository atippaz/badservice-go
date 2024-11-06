package databases

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Database interface {
	Connect() *mongo.Database
}
