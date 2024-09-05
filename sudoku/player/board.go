package player

import "../board"

func (p *Player) NewBoard() {
	p.board = board.Create(p.difficulty)
}

func (p *Player) FinishBoard() {
	if p.board.BoardDone() {
		p.AddWin()
		mistakes := p.board.GetMistakes()
		if mistakes == 0 {
			p.AddPerfectWin()
		}
		points := (100 - (mistakes * 5)) * p.difficulty
		p.AddPoints(points)
	} else {
		p.AddLoss()
	}
	p.board = nil
}
