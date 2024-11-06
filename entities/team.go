// Team represents the structure of teamSchema from Mongoose
package entities

type Team struct {
	Member []string `bson:"member"` // Array of strings
	Order  int      `bson:"order"`  // Integer field
}
