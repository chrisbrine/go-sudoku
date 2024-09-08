package game

type LastMoveType struct {
	Row int `json:"Row"`
	Col int `json:"Col"`
	Num int `json:"Num"`
	Type string `json:"Type"`
}

type GameResult struct {
	Board [9][9]int `json:"Board"`
	// PlayerBoard [9][9]int `json:"PlayerBoard"`
	Hints [9][9][9]bool `json:"Hints"`
	NumbersLeft [9]int `json:"NumbersLeft"`
	Mistakes int `json:"Mistakes"`
	Playing bool `json:"Playing"`
	InGame bool `json:"InGame"`
	GameStatus int `json:"GameStatus"`
	LastMove LastMoveType `json:"LastMove"`
	Username string `json:"Username"`
	Name string `json:"Name"`
	Wins int `json:"Wins"`
	Losses int `json:"Losses"`
	Points int `json:"Points"`
	PerfectWins int `json:"PerfectWins"`
	Difficulty int `json:"Difficulty"`
	Success bool `json:"Success"`
}

type Leader struct {
	Username string `json:"Username"`
	Name string `json:"Name"`
	Wins int `json:"Wins"`
	Losses int `json:"Losses"`
	Points int `json:"Points"`
	PerfectWins int `json:"PerfectWins"`
}
