package player

import (
	"github.com/chrisbrine/go-sudoku/sudoku/board"
	"github.com/google/uuid"
)

func (p *Player) SetDifficulty(difficulty int) {
	// if difficulty less than 1 or greater than 3 then set it to 1
	if difficulty < 1 || difficulty > 3 {
		difficulty = 1
	} else {
		p.Difficulty = difficulty
	}
}

func (p *Player) SetGame(board *board.Board) {
	p.Board = board
}

func (p *Player) GetGame() *board.Board {
	return p.Board
}

func (p *Player) GetDifficulty() int {
	return p.Difficulty
}

func (p *Player) GetWins() int {
	return p.Wins
}

func (p *Player) GetLosses() int {
	return p.Losses
}

func (p *Player) GetPerfectWins() int {
	return p.PerfectWins
}

func (p *Player) GetPoints() int {
	return p.Points
}

func (p *Player) AddWin() {
	p.Wins++
}

func (p *Player) AddLoss() {
	p.Losses++
}

func (p *Player) AddPerfectWin() {
	p.PerfectWins++
}

func (p *Player) AddPoints(points int) {
	p.Points += points
}

func (p *Player) SetupPlayer() {
	p.Wins = 0
	p.Losses = 0
	p.Points = 0
	p.SetDifficulty(1)
}

func (p *Player) SetID(id string) {
	// if no string passed, generate id instead
	if id == "" {
		p.Id = GenerateID()
	} else {
		p.Id = id
	}
}

func (p *Player) GetID() string {
	return p.Id
}

/* handle random ID generation, produce as UUID */

func GenerateID() string {
	// generate a random UUID
	id := uuid.New()
	return id.String()
}