package game

type GameResult struct {
	Board [9][9]int `json:"Board"`
	PlayerBoard [9][9]int `json:"PlayerBoard"`
	Hints [9][9][9]bool `json:"Hints"`
	NumbersLeft [9]int `json:"NumbersLeft"`
	Mistakes int `json:"Mistakes"`
	Playing bool `json:"Playing"`
	Username string `json:"Username"`
	Name string `json:"Name"`
	Wins int `json:"Wins"`
	Losses int `json:"Losses"`
	Points int `json:"Points"`
	PerfectWins int `json:"PerfectWins"`
	Difficulty int `json:"Difficulty"`
	Success bool `json:"Success"`
}
