package sql

import (
	"database/sql"
)

type DBData struct {
	db *sql.DB
}

// type DBGame struct {
// 	playerID string
// 	data string
// }

type DBPlayer struct {
	db *DBData
	Id int
	Username string
	Name string
	Password string
	PerfectWins int
	Wins int
	Losses int
	Points int
	Difficulty int
	Token string
}
