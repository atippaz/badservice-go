// Team represents the structure of teamSchema from Mongoose
package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Member []string `bson:"member"` // Array of strings
	Order  int      `bson:"order"`  // Integer field
}
