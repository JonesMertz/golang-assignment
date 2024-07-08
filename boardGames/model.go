package boardGames

type (
	BoardGame struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		MinPlayers  int    `json:"minPlayers"`
		MaxPlayers  int    `json:"maxPlayers"`
		PlayTime    int    `json:"playTime"`
		Age         int    `json:"age"`
		Description string `json:"description"`
		Price       int    `json:"price"`
	}
)

func GetDummyBoardGames() []BoardGame {
	return []BoardGame{
		{Name: "Catan", MinPlayers: 3, MaxPlayers: 4, PlayTime: 60, Age: 10, Description: "A fun game", Price: 50},
		{Name: "Monopoly", MinPlayers: 2, MaxPlayers: 6, PlayTime: 90, Age: 8, Description: "A boring game", Price: 20},
		{Name: "Risk", MinPlayers: 2, MaxPlayers: 6, PlayTime: 120, Age: 10, Description: "A command and conquer game", Price: 40},
		{Name: "Clue", MinPlayers: 3, MaxPlayers: 6, PlayTime: 45, Age: 8, Description: "A detective game", Price: 30},
		{Name: "Scrabble", MinPlayers: 2, MaxPlayers: 4, PlayTime: 90, Age: 10, Description: "A word game", Price: 25},
		{Name: "Ticket to Ride", MinPlayers: 2, MaxPlayers: 5, PlayTime: 60, Age: 8, Description: "A train adventure game", Price: 45},
		{Name: "Pandemic", MinPlayers: 2, MaxPlayers: 4, PlayTime: 45, Age: 10, Description: "A cooperative game", Price: 35},
		{Name: "Codenames", MinPlayers: 2, MaxPlayers: 8, PlayTime: 15, Age: 14, Description: "A word association game", Price: 20},
		{Name: "Splendor", MinPlayers: 2, MaxPlayers: 4, PlayTime: 30, Age: 10, Description: "A gem-collecting game", Price: 40},
		{Name: "Azul", MinPlayers: 2, MaxPlayers: 4, PlayTime: 45, Age: 8, Description: "A tile-placement game", Price: 30},
	}
}
