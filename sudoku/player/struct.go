package player

import sudokuBoard "../board"

type Player struct {
	id string
	board *sudokuBoard.Board
	perfectWins int
	wins int
	losses int
	points int
	difficulty int
}
