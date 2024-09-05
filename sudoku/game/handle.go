package game

import (
	"math/rand"

	"github.com/chrisbrine/go-sudoku/sudoku/player"
)

func (g *Game) SetKey(key string) {
	g.key = key
}

func (g *Game) GetKey() string {
	return g.key
}

func (g *Game) RandomKey() string {
	/* Create a random string encoded key tht is 24-30 characters long */

	// Create a random string
	keyChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"
	// Create a random number between 24 and 30
	length := 24 + rand.Intn(7)

	// Create a random key using the keyChars
	g.key = ""
	for i := 0; i < length; i++ {
		g.key += string(keyChars[rand.Intn(len(keyChars))])
	}

	return g.key
}

func (g *Game) SetPlayer(p *player.Player) {
	g.player = *p
}

func (g *Game) GetPlayer() *player.Player {
	return &g.player
}

func (g *Games) AddPlayer(p *player.Player) {
	// Create a new game for the player
	game := &Game{}
	game.SetPlayer(p)
	game.RandomKey()
	// g.games[game.key] = game
}

// func (g *Games)