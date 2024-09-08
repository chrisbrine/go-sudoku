package sql

import (
	"fmt"
	"math/rand"

	"github.com/chrisbrine/go-sudoku/sudoku/player"

	"golang.org/x/crypto/bcrypt"
)

func (d *DBData) CreatePlayerTable() error {
	_, err := d.db.Exec("CREATE TABLE IF NOT EXISTS players (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT NOT NULL UNIQUE, name TEXT, password TEXT, perfectWins INTEGER, wins INTEGER, losses INTEGER, points INTEGER, difficulty INTEGER, token TEXT UNIQUE)")
	if err != nil {
		return err
	}

	return nil
}

func (d *DBData) GetPlayer(username string) (DBPlayer, error) {
	var player DBPlayer
	err := d.db.QueryRow("SELECT * FROM players WHERE username = ?", username).Scan(
		&player.Id,
		&player.Username,
		&player.Name,
		&player.Password,
		&player.PerfectWins,
		&player.Wins,
		&player.Losses,
		&player.Points,
		&player.Difficulty,
		&player.Token,
	)
	if err != nil {
		fmt.Println("Error getting player", username, ":", err)
		return player, err
	}

	player.db = d

	return player, nil
}

func (d *DBData) AddPlayer(player *DBPlayer) error {
	newPassword, passErr := hashPassword(player.Password)
	if passErr != nil {
		return passErr
	}
	player.Password = newPassword

	_, dbErr := d.db.Exec("INSERT INTO players (username, name, password, perfectWins, wins, losses, points, difficulty) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", player.Username, player.Name, player.Password, player.PerfectWins, player.Wins, player.Losses, player.Points, player.Difficulty)
	if dbErr != nil {
		return dbErr
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (p *DBPlayer) GetPlayerData() player.Player {
	game, gameErr := p.db.GetGame(p.Id)
	if gameErr != nil {
		game = nil
	}
	return player.Player{
		Id:          p.Id,
		Username:    p.Username,
		Name:        p.Name,
		PerfectWins: p.PerfectWins,
		Wins:        p.Wins,
		Losses:      p.Losses,
		Points:      p.Points,
		Difficulty:  p.Difficulty,
		Board:       game,
	}
}

func (p *DBPlayer) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (p *DBPlayer) GetUsername() string {
	return p.Username
}

func (p *DBPlayer) GetName() string {
	return p.Name
}

func (d *DBData) UpdatePlayer(player *player.Player) error {
	_, err := d.db.Exec("UPDATE players SET perfectWins = ?, wins = ?, losses = ?, points = ?, difficulty = ? WHERE id = ?", player.GetPerfectWins(), player.GetWins(), player.GetLosses(), player.GetPoints(), player.GetDifficulty(), player.Id)
	if err != nil {
		return err
	}

	return nil
}

func (d *DBData) DeletePlayer(player *DBPlayer) error {
	_, err := d.db.Exec("DELETE FROM players WHERE id = ?", player.Id)
	if err != nil {
		return err
	}

	return nil
}

func (d *DBData) ChangePassword(player *DBPlayer, password string) error {
	newPassword, passErr := hashPassword(password)
	if passErr != nil {
		return passErr
	}

	_, dbErr := d.db.Exec("UPDATE players SET password = ? WHERE id = ?", newPassword, player.Id)
	if dbErr != nil {
		return dbErr
	}

	return nil
}

func (d *DBData) ChangeUsername(player *DBPlayer, username string) error {
	_, err := d.db.Exec("UPDATE players SET username = ? WHERE id = ?", username, player.Id)
	if err != nil {
		return err
	}

	return nil
}

func (d *DBData) ChangeName(player *DBPlayer, name string) error {
	_, err := d.db.Exec("UPDATE players SET name = ? WHERE id = ?", name, player.Id)
	if err != nil {
		return err
	}

	return nil
}

func (d *DBData) UpdateDifficulty(player *DBPlayer, difficulty int) error {
	_, err := d.db.Exec("UPDATE players SET difficulty = ? WHERE id = ?", difficulty, player.Id)
	if err != nil {
		return err
	}

	return nil
}

func (db *DBData) Login(username string, password string) (bool, error) {
	var count int
	dbErr := db.db.QueryRow("SELECT COUNT(*) FROM players WHERE username = ?", username).Scan(&count)
	if dbErr != nil {
		return false, dbErr
	}

	if count == 0 {
		return false, nil
	}

	player, playerErr := db.GetPlayer(username)
	if playerErr != nil {
		return false, playerErr
	}

	if player.CheckPasswordHash(password, player.Password) {
		_, tokenErr := db.CreateToken(&player)
		if tokenErr != nil {
			return false, tokenErr
		}

		return true, nil
	}

	return false, nil
}

func (db *DBData) CreateToken(player *DBPlayer) (string, error) {
	// create a random token and then set it on the player in the database and the player struct

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+1234567890"
	token := ""

	for token == "" {
		b := make([]byte, 64)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}

		token = string(b)

		// make sure it doesn't exist in the database
		var count int

		dbErr := db.db.QueryRow("SELECT COUNT(*) FROM players WHERE token = ?", token).Scan(&count)
		if dbErr != nil {
			return "", dbErr
		}

		if count == 0 {
			break
		}

		token = ""
	}

	_, dbErr := db.db.Exec("UPDATE players SET token = ? WHERE username = ?", token, player.Username)
	if dbErr != nil {
		return "", dbErr
	}

	player.Token = token

	return token, nil
}

func (db *DBData) GetPlayerFromToken(token string) (DBPlayer, error) {
	var player DBPlayer
	err := db.db.QueryRow("SELECT * FROM players WHERE token = ?", token).Scan(&player.Id, &player.Username, &player.Name, &player.Password, &player.PerfectWins, &player.Wins, &player.Losses, &player.Points, &player.Difficulty, &player.Token)
	if err != nil {
		return player, err
	}

	player.db = db

	return player, nil
}

func (db *DBData) ConfirmToken(token string, username string) (bool, error) {
	var count int
	sqlRow := db.db.QueryRow("SELECT COUNT(*) FROM players WHERE token = ? AND username = ?", token, username)
	dbErr := sqlRow.Scan(&count)
	if dbErr != nil {
		return false, dbErr
	}

	if count == 1 {
		return true, nil
	}

	return false, nil
}

func (db *DBData) DeleteToken(player *DBPlayer) error {
	_, err := db.db.Exec("UPDATE players SET token = ? WHERE id = ?", "", player.Id)
	if err != nil {
		return err
	}

	return nil
}

func (db *DBData) GetToken(player *DBPlayer) string {
	return player.Token
}
