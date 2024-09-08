package player

import "github.com/chrisbrine/go-sudoku/sudoku/board"

type Player struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`	
	Board *board.Board `json:"board"`
	PerfectWins int `json:"perfectWins"`
	Wins int `json:"wins"`
	Losses int `json:"losses"`
	Points int `json:"points"`
	Difficulty int `json:"difficulty"`
}
