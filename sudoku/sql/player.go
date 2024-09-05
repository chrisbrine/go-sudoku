package sql

import (
	"github.com/chrisbrine/go-sudoku/sudoku/player"
	"golang.org/x/crypto/bcrypt"
)

func (d *DB) CreatePlayerTable() error {
	_, err := d.db.Exec("CREATE TABLE IF NOT EXISTS players (id TEXT PRIMARY KEY, username TEXT NOT NULL UNIQUE, name TEXT, password TEXT, perfectWins INTEGER, wins INTEGER, losses INTEGER, points INTEGER, difficulty INTEGER)")
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetPlayer(username string) (DBPlayer, error) {
	var player DBPlayer
	err := d.db.QueryRow("SELECT * FROM players WHERE username = ?", username).Scan(&player.id, &player.username, &player.name, &player.password, &player.perfectWins, &player.wins, &player.losses, &player.points, &player.difficulty)
	if err != nil {
		return player, err
	}

	return player, nil
}

func (d *DB) AddPlayer(player *DBPlayer) error {
	newPassword, passErr := hashPassword(player.password)
	if passErr != nil {
		return passErr
	}
	player.password = newPassword

	_, dbErr := d.db.Exec("INSERT INTO players (id, username, name, password, perfectWins, wins, losses, points, difficulty) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", player.id, player.username, player.name, player.password, player.perfectWins, player.wins, player.losses, player.points, player.difficulty)
	if dbErr != nil {
		return dbErr
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (p *DBPlayer) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (d *DB) UpdatePlayer(player *player.Player) error {
	_, err := d.db.Exec("UPDATE players SET perfectWins = ?, wins = ?, losses = ?, points = ?, difficulty = ? WHERE id = ?", player.GetPerfectWins(), player.GetWins(), player.GetLosses(), player.GetPoints(), player.GetDifficulty(), player.GetID())
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) DeletePlayer(player *DBPlayer) error {
	_, err := d.db.Exec("DELETE FROM players WHERE id = ?", player.id)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) ChangePassword(player *DBPlayer, password string) error {
	newPassword, passErr := hashPassword(password)
	if passErr != nil {
		return passErr
	}

	_, dbErr := d.db.Exec("UPDATE players SET password = ? WHERE id = ?", newPassword, player.id)
	if dbErr != nil {
		return dbErr
	}

	return nil
}

func (d *DB) ChangeUsername(player *DBPlayer, username string) error {
	_, err := d.db.Exec("UPDATE players SET username = ? WHERE id = ?", username, player.id)
	if err != nil {
		return err
	}

	return nil
}
