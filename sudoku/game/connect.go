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

func Register(db *sudokuSQL.DBData, username string, password string, name string) (string, error) {
	// check if database is connected
	if db == nil {
		// do not use an sql error
		return "", errors.New("database not connected")
	}

	// check if username and password are correct
	if username == "" || password == "" || name == "" {
		return "", errors.New("blank user credentials")
	}

	// create a DBPlayer overlay
	DBPlayer := sudokuSQL.DBPlayer{
		Username: username,
		Password: password,
		Name:     name,
		Wins:		 0,
		Losses:	 0,
		PerfectWins: 0,
		Points: 0,
		Difficulty: 1,
		Token: 	"",
	}

	// register the player
	playerErr := db.AddPlayer(&DBPlayer)
	if playerErr != nil {
		return "", playerErr
	}

	// create token
	token, tokenErr := db.CreateToken(&DBPlayer)
	if tokenErr != nil {
		return "", tokenErr
	}

	return token, nil
}

func ChangePassword(db *sudokuSQL.DBData, token string, oldPassword string, newPassword string) (bool, error) {
	// check if database is connected
	if db == nil {
		// do not use an sql error
		return false, errors.New("database not connected")
	}

	// check if token is correct
	if token == "" {
		// non sql error for wrong login credentials
		return false, errors.New("wrong token")
	}

	player, playerError := db.GetPlayerFromToken(token)
	if playerError != nil {
		return false, playerError
	}

	// verify the old password
	if !player.CheckPasswordHash(oldPassword, player.Password) {
		return false, errors.New("wrong password")
	}

	// change the password
	passwordError := db.ChangePassword(&player, newPassword)
	if passwordError != nil {
		return false, passwordError
	}

	return true, nil
}

func ChangeUserName(db *sudokuSQL.DBData, token string, newUserName string) (bool, error) {
	// check if database is connected
	if db == nil {
		// do not use an sql error
		return false, errors.New("database not connected")
	}

	// check if token is correct
	if token == "" {
		// non sql error for wrong login credentials
		return false, errors.New("wrong token")
	}

	player, playerError := db.GetPlayerFromToken(token)
	if playerError != nil {
		return false, playerError
	}

	// change the username
	usernameError := db.ChangeUsername(&player, newUserName)
	if usernameError != nil {
		return false, usernameError
	}

	return true, nil
}

func ChangeName(db *sudokuSQL.DBData, token string, newName string) (bool, error) {
	// check if database is connected
	if db == nil {
		// do not use an sql error
		return false, errors.New("database not connected")
	}

	// check if token is correct
	if token == "" {
		// non sql error for wrong login credentials
		return false, errors.New("wrong token")
	}

	player, playerError := db.GetPlayerFromToken(token)
	if playerError != nil {
		return false, playerError
	}

	// change the name
	nameError := db.ChangeName(&player, newName)
	if nameError != nil {
		return false, nameError
	}

	return true, nil
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
