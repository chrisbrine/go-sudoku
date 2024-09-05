package game

import (
	"errors"

	sudokuSQL "github.com/chrisbrine/go-sudoku/sudoku/sql"
)

func Connect(path string) (*sudokuSQL.DBData, error) {
	db, err := sudokuSQL.Connect(path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Login(db *sudokuSQL.DBData, username string, password string) (string, error) {
	// check if database is connected
	if db == nil {
		// do not use an sql error
		return "", errors.New("database not connected")
	}

	// check if username and password are correct
	if username == "" || password == "" {
		return "", errors.New("wrong login credentials")
	}

	success, loginError := db.Login(username, password)
	if loginError != nil {
		return "", loginError
	}
	if !success {
		return "", errors.New("wrong login credentials")
	}

	// get the token
	player, playerError := db.GetPlayer(username)
	if playerError != nil {
		return "", playerError
	}

	token := db.GetToken(&player)

	if token == "" {
		return "", errors.New("token not found")
	}

	return token, nil
}

func Logout(db *sudokuSQL.DBData, token string) error {
	// check if database is connected
	if db == nil {
		// do not use an sql error
		return errors.New("database not connected")
	}

	// check if token is correct
	if token == "" {
		// non sql error for wrong login credentials
		return errors.New("wrong token")
	}

	player, playerError := db.GetPlayerFromToken(token)
	if playerError != nil {
		return playerError
	}

	// delete the token
	tokenError := db.DeleteToken(&player)
	if tokenError != nil {
		return tokenError
	}

	return nil
}