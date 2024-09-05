package player

import "github.com/chrisbrine/go-sudoku/sudoku/board"

type Player struct {
	Id string `json:"id"`
	Board *board.Board `json:"board"`
	PerfectWins int `json:"perfectWins"`
	Wins int `json:"wins"`
	Losses int `json:"losses"`
	Points int `json:"points"`
	Difficulty int `json:"difficulty"`
}
