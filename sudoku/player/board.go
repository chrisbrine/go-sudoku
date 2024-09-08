package player

import "github.com/chrisbrine/go-sudoku/sudoku/board"

func (p *Player) NewBoard() {
	p.Board = board.Create(p.Difficulty)
}

func (p *Player) FinishBoard() {
	if p.Board.BoardDone() {
		p.AddWin()
		mistakes := p.Board.GetMistakes()
		if mistakes == 0 {
			p.AddPerfectWin()
		}
		points := (100 - (mistakes * 5)) * p.Difficulty
		p.AddPoints(points)
	} else if p.Board.GetMistakes() > 20 {
		points := -25
		if p.Points - points < 0 {
			points = -p.Points
		}
		p.AddPoints(points)
		p.AddLoss()
	}
	p.Board = nil
}
