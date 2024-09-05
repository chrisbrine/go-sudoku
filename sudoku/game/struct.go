package game

import "player"

type Game struct {
	key string
	player player.Player
}

type Games struct {
	games map[string]Game
}
