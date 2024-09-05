package sql

import "github.com/chrisbrine/go-sudoku/sudoku/board"

func (db *DBData) CreateGameTable() error {
	_, err := db.db.Exec("CREATE TABLE IF NOT EXISTS games (player_id TEXT NOT NULL UNIQUE, data TEXT NOT NULL UNIQUE)")
	if err != nil {
		return err
	}

	return nil
}

func (db *DBData) AddGame(playerID string, b *board.Board) error {
	// get a json string of the board
	data, boardErr := b.ToJson()
	if boardErr != nil {
		return boardErr
	}

	_, dbErr := db.db.Exec("INSERT INTO games (player_id, data) VALUES ($1, $2)", playerID, data)
	if dbErr != nil {
		return dbErr
	}

	return nil
}

func (db *DBData) DeleteGame(playerID string) error {
	_, dbErr := db.db.Exec("DELETE FROM games WHERE player_id = $1", playerID)
	if dbErr != nil {
		return dbErr
	}

	return nil
}

func (db *DBData) GetGame(playerID string) (*board.Board, error) {
	var data string
	dbErr := db.db.QueryRow("SELECT data FROM games WHERE player_id = $1", playerID).Scan(&data)
	if dbErr != nil {
		return nil, dbErr
	}

	b, boardErr := board.CreateFromJson(data)
	if boardErr != nil {
		return nil, boardErr
	}

	return b, nil
}
