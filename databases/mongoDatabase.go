package databases

import (
	"context"
	"sync"

	Config "bad-service-go/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDatabase struct {
	*mongo.Database
}

var (
	mongoDatabaseInstance *mongoDatabase
	once                  sync.Once
)

// Connect returns the MongoDB client instance
func (db *mongoDatabase) Connect() *mongo.Database {
	return db.Database
}

// NewMongoDatabase initializes the MongoDB connection as a singleton instance
func NewMongoDatabase(conf *Config.Database) (*mongoDatabase, error) {
	var initError error

	once.Do(func() {
		uri := conf.Url
		databaseName := conf.DatabaseName

		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			initError = err
			return
		}

		// ตรวจสอบการเชื่อมต่อ
		if err := client.Ping(context.TODO(), nil); err != nil {
			initError = err
			return
		}
		database := client.Database(databaseName)
		mongoDatabaseInstance = &mongoDatabase{database}
	})

	return mongoDatabaseInstance, initError
}
