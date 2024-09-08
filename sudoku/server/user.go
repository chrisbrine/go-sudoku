package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/chrisbrine/go-sudoku/sudoku/game"
	"github.com/chrisbrine/go-sudoku/sudoku/sql"
)

func UserLogin(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string) {
	data := GetFromBody(r, []string{"username", "password"})
	username := data["username"]
	password := data["password"]

	fmt.Println("Logging in user", username)

	token, err := game.Login(db, username, password)
	if err != nil {
		fmt.Println("Error logging in user", username, ":", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	RespondToken(w, token)
}

func UserRegister(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string) {
	data := GetFromBody(r, []string{"username", "password", "name"})
	username := data["username"]
	password := data["password"]
	name := data["name"]

	fmt.Println("Registering user", username)

	token, err := game.Register(db, username, password, name)
	if err != nil {
		fmt.Println("Error registering user", username, ":", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	RespondToken(w, token)
}

func UserLogout(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string) {
	token := params["token"]

	err := game.Logout(db, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	RespondSuccess(w, true)
}

func UserUpdateUsername(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string) {
	token := GetToken(r)
	newUsername := GetOneFromBody(r, "newUsername")

	success, err := game.ChangeUserName(db, token, newUsername)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	RespondSuccess(w, success)
}

func UserUpdatePassword(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string) {
	token := GetToken(r)
	data := GetFromBody(r, []string{"oldPassword", "newPassword"})
	oldPassword := data["oldPassword"]
	newPassword := data["newPassword"]

	success, err := game.ChangePassword(db, token, oldPassword, newPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	RespondSuccess(w, success)
}

func UserUpdateName(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string) {
	token := GetToken(r)
	newName := GetOneFromBody(r, "newName")

	success, err := game.ChangeName(db, token, newName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	RespondSuccess(w, success)
}

func UserUpdateDifficulty(db *sql.DBData, w http.ResponseWriter, r *http.Request, params map[string]string) {
	token := GetToken(r)
	newDifficulty := params["difficulty"]

	intDifficulty, err := strconv.Atoi(newDifficulty)
	if err != nil {
		fmt.Println("Error converting difficulty to int:", err, newDifficulty)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	gameData, err := game.ChangeDifficulty(db, token, intDifficulty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	RespondGame(w, gameData)
}

func AddUserMethods(db *sql.DBData) {
	HandlePOST(db, "/api/login", UserLogin, false)
	HandlePOST(db, "/api/register", UserRegister, false)
	HandlePOST(db, "/api/logout", UserLogout, true)
	HandlePOST(db, "/api/update/username", UserUpdateUsername, true)
	HandlePOST(db, "/api/update/password", UserUpdatePassword, true)
	HandlePOST(db, "/api/update/name", UserUpdateName, true)
	HandleGET(db, "/api/update/difficulty/{difficulty}", UserUpdateDifficulty, true)
}