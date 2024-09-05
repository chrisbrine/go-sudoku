package game

import (
	"github.com/chrisbrine/go-sudoku/sudoku/player"
	"github.com/chrisbrine/go-sudoku/sudoku/sql"
)

func GetData(db *sql.DBData, token string) (sql.DBPlayer, player.Player, error) {
	// get player data
	playerData, err := db.GetPlayerFromToken(token)
	if err != nil {
		return sql.DBPlayer{}, player.Player{}, err
	}
	player := playerData.GetPlayerData()

	// return player data
	return playerData, player, nil
}

func NewGame(db *sql.DBData, token string) (string, error) {
	// get player data
	playerData, playerInfo, playerErr := GetData(db, token)
	if playerErr != nil {
		return "", playerErr
	}

	// create a new game
	playerInfo.NewBoard()

	// write board to database
	dbErr := db.AddGame(playerInfo.Id, playerInfo.GetGame())
	if dbErr != nil {
		return "", dbErr
	}

	// return the result as json
	return HandleResult(&playerInfo, &playerData, true)
}

func PickNumber(db *sql.DBData, token string, row int, col int, num int) (string, error) {
	// get player data
	playerData, playerInfo, playerErr := GetData(db, token)
	if playerErr != nil {
		return "", playerErr
	}

	// get board
	board, boardErr := GetBoard(db, playerInfo)
	if boardErr != nil {
		return "", boardErr
	}

	// set the number
	success := board.SetMove(row, col, num)

	// write board to database
	dbErr := db.AddGame(playerInfo.Id, board)
	if dbErr != nil {
		return "", dbErr
	}

	// return the result as json
	return HandleResult(&playerInfo, &playerData,success)
}

func PickHint(db *sql.DBData, token string, row int, col int, num int, remove bool) (string, error) {
	// get player data
	playerData, playerInfo, playerErr := GetData(db, token)
	if playerErr != nil {
		return "", playerErr
	}

	// get board
	board, boardErr := GetBoard(db, playerInfo)
	if boardErr != nil {
		return "", boardErr
	}

	// set the hint
	var success bool
	if remove {
		success = true
		board.RemoveHint(row, col, num)
	} else {
		success = board.SetHint(row, col, num)
	}

	// write board to database
	dbErr := db.AddGame(playerInfo.Id, board)
	if dbErr != nil {
		return "", dbErr
	}

	// return the result as json
	return HandleResult(&playerInfo, &playerData, success)
}

func GetCurrent(db *sql.DBData, token string) (string, error) {
	// get player data
	playerData, playerInfo, playerErr := GetData(db, token)
	if playerErr != nil {
		return "", playerErr
	}

	// return the result as json
	return HandleResult(&playerInfo, &playerData, true)
}

func QuitGame(db *sql.DBData, token string) (string, error) {
	// get player data
	playerData, playerInfo, playerErr := GetData(db, token)
	if playerErr != nil {
		return "", playerErr
	}

	// get board
	board, boardErr := GetBoard(db, playerInfo)
	if boardErr != nil {
		return "", boardErr
	}

	board.QuitGame()

	// delete board from database
	dbErr := db.DeleteGame(playerInfo.Id)
	if dbErr != nil {
		return "", dbErr
	}

	// return the result as json
	return HandleResult(&playerInfo, &playerData, true)
}
