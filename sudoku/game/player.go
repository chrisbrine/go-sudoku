package game

import (
	"github.com/chrisbrine/go-sudoku/sudoku/board"
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

func CheckEndGame(db *sql.DBData, playerInfo *player.Player, board *board.Board) (bool, error) {
		// check if the game is over
	if board.BoardDone() {
		playerInfo.FinishBoard()

		// write playerInfo to database
		dbErr := db.UpdatePlayer(playerInfo)
		if dbErr != nil {
			return false, dbErr
		}

		// delete board from database
		dbErr = db.DeleteGame(playerInfo.Id)
		if dbErr != nil {
			return false, dbErr
		}

		return true, nil
	}

	return false, nil
}

func ChangeDifficulty(db *sql.DBData, token string, difficulty int) (string, error) {
	// get player data
	playerData, playerInfo, playerErr := GetData(db, token)
	if playerErr != nil {
		return "", playerErr
	}

	// change difficulty
	db.UpdateDifficulty(&playerData, difficulty)

	// write playerInfo to database
	dbErr := db.UpdatePlayer(&playerInfo)
	if dbErr != nil {
		return "", dbErr
	}

	// return the result as json
	return HandleResult(&playerInfo, &playerData, true)
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
	playerInfo.Board = board

	// write board to database
	dbErr := db.AddGame(playerInfo.Id, playerInfo.Board)
	if dbErr != nil {
		return "", dbErr
	}

	// check if the game is over
	_, endGameErr := CheckEndGame(db, &playerInfo, playerInfo.Board)
	if endGameErr != nil {
		return "", endGameErr
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

	playerInfo.Board = board

	// set the hint
	var success bool
	if remove {
		success = true
		playerInfo.Board.RemoveHint(row, col, num)
	} else {
		success = playerInfo.Board.SetHint(row, col, num)
	}

	// write board to database
	dbErr := db.AddGame(playerInfo.Id, playerInfo.Board)
	if dbErr != nil {
		return "", dbErr
	}

	// check if the game is over
	_, endGameErr := CheckEndGame(db, &playerInfo, playerInfo.Board)
	if endGameErr != nil {
		return "", endGameErr
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
