package sql

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
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
