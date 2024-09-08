package server

import (
	"fmt"
	"net/http"

	"github.com/chrisbrine/go-sudoku/sudoku/game"
	"github.com/chrisbrine/go-sudoku/sudoku/sql"
)

func GetGame(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string) {
	game, err := game.GetCurrent(db, params["token"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	RespondGame(w, game)
}

func NewGame(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string) {
	game, err := game.NewGame(db, params["token"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	RespondGame(w, game)
}

func QuitGame(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string) {
	game, err := game.QuitGame(db, params["token"])
	if err != nil {
		fmt.Println("Error quitting game: ", err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		// return
	}

	RespondGame(w, game)
}

func SetMove(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string) {
	row, col, num, err := GetRowColNum(params)
	row--
	col--
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	game, err := game.PickNumber(db, params["token"], row, col, num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	RespondGame(w, game)
}

func HandleHint(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string, remove bool) {
	row, col, num, err := GetRowColNum(params)
	row--
	col--
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	game, err := game.PickHint(db, params["token"], row, col, num, remove)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	RespondGame(w, game)
}

func SetHint(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string) {
	HandleHint(db, w, r, params, false)
}

func RemoveHint(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string) {
	HandleHint(db, w, r, params, true)
}

func GetLeaderboard(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string) {
	leaderboard, err := game.GetLeaderboard(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	RespondGame(w, leaderboard)
}

func AddGameMethods(db *sql.DBData) {
	HandleGET(db,  "/api/game", GetGame, true)
	HandleGET(db, "/api/game/new", NewGame, true)
	HandleGET(db, "/api/game/quit", QuitGame, true)
	HandleGET(db, "/api/game/leaderboard", GetLeaderboard, true)
	HandlePOST(db, "/api/game/move/{row}/{col}/{num}", SetMove, true)
	HandlePOST(db, "/api/game/hint/{row}/{col}/{num}", SetHint, true)
	HandlePOST(db, "/api/game/hintRemove/{row}/{col}/{num}", RemoveHint, true)
}