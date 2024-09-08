package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/chrisbrine/go-sudoku/sudoku/sql"
)

func StartServer(db *sql.DBData, port int) {
	// ALWAYS get the auth token from the header

	// from ANY METHOD /api/health to check if the server is running
	// from GET /api/game get the current game
	// from POST /api/game/new create a new game
	// from POST /api/game/move/{row}/{col}/{num} set a number in the game
	// from POST /api/game/hint/{row}/{col}/{num} set a hint in the game
	// from POST /api/game/hintRemove/{row}/{col}/{num} remove a hint in the game
	// from POST /api/login (posts username and password) and send back the token
	// from POST /api/register (posts username, password, and name) and send back the token
	// from POST /api/logout (posts token) and remove the token from the database
	// from POST /api/update/password (posts new password and old password) and change the password in the database
	// from POST /api/update/name (posts new name) and change the name in the database
	// from POST /api/update/difficulty (posts new difficulty) and change the difficulty in the database
	// FROM POST /api/update/username (posts new username) and change the username in the database

	// add all the methods to the server
	AddGameMethods(db)
	AddUserMethods(db)
	AddMiscMethods(db)

	// serve the public directory
	http.Handle("/", http.FileServer(http.Dir("./build")))

	// start server
	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}