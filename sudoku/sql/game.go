package sql

import (
	"github.com/chrisbrine/go-sudoku/sudoku/board"
)

func (db *DB) CreateGameTable() error {
	_, err := db.db.Exec("CREATE TABLE IF NOT EXISTS games (player_id TEXT NOT NULL UNIQUE, data TEXT NOT NULL UNIQUE)")
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) AddGame(playerID string, b *board.Board) error {
	// get a json string of the board
	data = b.toJson()
	// convert data to string
	dataString := string(data)

	_, dbErr := db.db.Exec("INSERT INTO games (player_id, data) VALUES ($1, $2)", playerID, dataString)
	if dbErr != nil {
		return dbErr
	}

	return nil
}
