package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/chrisbrine/go-sudoku/sudoku/game"
	"github.com/chrisbrine/go-sudoku/sudoku/sql"
)

func StartServer(db sql.DBData, port int) {
	// server a public directory from ../ to /

	http.Handle("/", http.FileServer(http.Dir("../public")))

	// ALWAYS get the auth token from the header

	// from GET /api/game get the current game
	// from GET /api/game/new create a new game
	// from POST /api/game/move/{row}/{col}/{num} set a number in the game
	// from POST /api/game/hint/{row}/{col}/{num} set a hint in the game
	// from POST /api/game/hintRemove/{row}/{col}/{num} remove a hint in the game

	http.HandleFunc("/api/game", func(w http.ResponseWriter, r *http.Request) {
		// make sure it is a GET request
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// get the token from the header
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "No token provided", http.StatusUnauthorized)
			return
		}

		// get the game from the database
		game, err := game.GetCurrent(&db, token)

		// if there is an error, return it
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// if there is a game then send it with a json header
		if game != "" {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(game))
			return
		}
	})

	http.HandleFunc("/api/game/new", func(w http.ResponseWriter, r *http.Request) {
		// make sure it is a GET request
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// get the token from the header
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "No token provided", http.StatusUnauthorized)
			return
		}

		// create a new game
		game, err := game.NewGame(&db, token)

		// if there is an error, return it
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// send the game with a json header
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(game))
	})

	http.HandleFunc("/api/game/move/", func(w http.ResponseWriter, r *http.Request) {
		// make sure it is a POST request
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// get the token from the header
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "No token provided", http.StatusUnauthorized)
			return
		}

		// get the row, col, and num from the url
		row, err := strconv.Atoi(r.URL.Path[13:14])
		if err != nil {
			http.Error(w, "Invalid row", http.StatusBadRequest)
			return
		}

		col, err := strconv.Atoi(r.URL.Path[15:16])
		if err != nil {
			http.Error(w, "Invalid col", http.StatusBadRequest)
			return
		}

		num, err := strconv.Atoi(r.URL.Path[17:18])
		if err != nil {
			http.Error(w, "Invalid num", http.StatusBadRequest)
			return
		}

		// set the number in the game
		game, err := game.PickNumber(&db, token, row, col, num)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// send the game with a json header
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(game))
	})

	http.HandleFunc("/api/game/hint/", func(w http.ResponseWriter, r *http.Request) {
		// make sure it is a POST request
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// get the token from the header
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "No token provided", http.StatusUnauthorized)
			return
		}

		// get the row, col, and num from the url
		row, err := strconv.Atoi(r.URL.Path[13:14])
		if err != nil {
			http.Error(w, "Invalid row", http.StatusBadRequest)
			return
		}

		col, err := strconv.Atoi(r.URL.Path[15:16])
		if err != nil {
			http.Error(w, "Invalid col", http.StatusBadRequest)
			return
		}

		num, err := strconv.Atoi(r.URL.Path[17:18])
		if err != nil {
			http.Error(w, "Invalid num", http.StatusBadRequest)
			return
		}

		// set the hint in the game
		game, err := game.PickHint(&db, token, row, col, num, false)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// send the game with a json header
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(game))
	})

	http.HandleFunc("/api/game/hintRemove/", func(w http.ResponseWriter, r *http.Request) {
		// make sure it is a POST request
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// get the token from the header
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "No token provided", http.StatusUnauthorized)
			return
		}

		// get the row, col, and num from the url
		row, err := strconv.Atoi(r.URL.Path[18:19])
		if err != nil {
			http.Error(w, "Invalid row", http.StatusBadRequest)
			return
		}

		col, err := strconv.Atoi(r.URL.Path[20:21])
		if err != nil {
			http.Error(w, "Invalid col", http.StatusBadRequest)
			return
		}

		num, err := strconv.Atoi(r.URL.Path[22:23])
		if err != nil {
			http.Error(w, "Invalid num", http.StatusBadRequest)
			return
		}

		// remove the hint in the game
		game, err := game.PickHint(&db, token, row, col, num, true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// send the game with a json header
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(game))
	})

	// start server
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}