package sql

import (
	"database/sql"
)

func Connect(path string) (*DBData, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return &DBData{db}, nil
}

func (d *DBData) Close() error {
	return d.db.Close()
}
