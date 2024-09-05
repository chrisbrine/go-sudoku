package player

import (
	sudokuBoard "../board"
	"github.com/google/uuid"
)

func (p *Player) SetDifficulty(difficulty int) {
	// if difficulty less than 1 or greater than 3 then set it to 1
	if difficulty < 1 || difficulty > 3 {
		difficulty = 1
	} else {
		p.difficulty = difficulty
	}
}

func (p *Player) SetGame(board *sudokuBoard.Board) {
	p.board = board
}

func (p *Player) GetGame() *sudokuBoard.Board {
	return p.board
}

func (p *Player) GetDifficulty() int {
	return p.difficulty
}

func (p *Player) GetWins() int {
	return p.wins
}

func (p *Player) GetLosses() int {
	return p.losses
}

func (p *Player) GetPerfectWins() int {
	return p.perfectWins
}

func (p *Player) GetPoints() int {
	return p.points
}

func (p *Player) AddWin() {
	p.wins++
}

func (p *Player) AddLoss() {
	p.losses++
}

func (p *Player) AddPerfectWin() {
	p.perfectWins++
}

func (p *Player) AddPoints(points int) {
	p.points += points
}

func (p *Player) SetupPlayer() {
	p.wins = 0
	p.losses = 0
	p.points = 0
	p.SetDifficulty(1)
}

func (p *Player) SetID(id string) {
	// if no string passed, generate id instead
	if id == "" {
		p.id = GenerateID()
	} else {
		p.id = id
	}
}

func (p *Player) GetID() string {
	return p.id
}

/* handle random ID generation, produce as UUID */

func GenerateID() string {
	// generate a random UUID
	id := uuid.New()
	return id.String()
}