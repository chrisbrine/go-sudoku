package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

var db *sql.DB

func Start(path string) {
	var err error
	db, err = sql.Connect(path)
	if err != nil {
		log.Fatal(err)
	}
}

func StartServer(port int) {
	// server a public directory from ../ to /

	http.Handle("/", http.FileServer(http.Dir("../public")))

	// ALWAYS get the auth token from the header

	// from GET /api/game get the current game
	// from GET /api/game/new create a new game
	// from POST /api/game/move/{row}/{col}/{num} set a number in the game
	// from POST /api/game/hint/{row}/{col}/{num} set a hint in the game
	// from POST /api/game/hintRemove/{row}/{col}/{num} remove a hint in the game

	http.HandleFunc("/api/game", func(w http.ResponseWriter, r *http.Request) {
		// get the token from the header
		token := r.Header.Get("Authorization")

		// c
		// run game.NewGame(db, token)

	log.Fatal(http.ListenAndServe(":8080", nil))
}