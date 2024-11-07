package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type TeamSet struct {
	Order int    `bson:"order"`
	Set   []Team `bson:"set"`
}
type MatchSet struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	SetName     string             `bson:"setName"`
	RoomID      string             `bson:"roomId"`
	AllTeam     []TeamSet          `bson:"allTeam"`
	CourtNumber int                `bson:"courtNumber"`
	TeamLimit   int                `bson:"teamLimit"`
	WinScore    int                `bson:"winScore"`
	WinStreak   int                `bson:"winStreak"`
}
