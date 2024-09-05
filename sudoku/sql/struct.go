package sql

import (
	"database/sql"
)

type DB struct {
	db *sql.DB
}

type DBGame struct {
	playerID string
	data string
}

type DBPlayer struct {
	id string
	username string
	name string
	password string
	perfectWins int
	wins int
	losses int
	points int
	difficulty int
}
