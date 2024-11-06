package entities

type Match struct {
	RoomID      string `bson:"roomId"`      // Room ID as a string
	SetName     string `bson:"setName"`     // Set name as a string
	AllTeam     []Team `bson:"allTeam"`     // Array of Team structs
	CourtNumber int    `bson:"courtNumber"` // Court number as an integer
	TeamLimit   int    `bson:"teamLimit"`   // Team limit as an integer
	WinScore    int    `bson:"winScore"`    // Winning score as an integer
	WinStreak   int    `bson:"winStreak"`   // Win streak as an integer
}
