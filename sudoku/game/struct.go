package game

import "github.com/chrisbrine/go-sudoku/sudoku/player"

type Game struct {
	key string
	player player.Player
}

type Games struct {
	games map[string]Game
}
