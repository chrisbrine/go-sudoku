package sudoku

import (
	"log"

	"github.com/chrisbrine/go-sudoku/sudoku/server"
	"github.com/chrisbrine/go-sudoku/sudoku/sql"
)

func Run(port int, path string) {
	// server.Start(path)
	// server.StartServer(port)
	db, err := sql.Start(path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server.StartServer(db, port)
}
