package game

import (
	"fmt"

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

func ChangeDifficulty(db *sql.DBData, token string, difficulty int) (string, error) {
	// get player data
	playerData, playerInfo, playerErr := GetData(db, token)
	if playerErr != nil {
		return "", playerErr
	}

	// change difficulty
	playerData.Difficulty = difficulty
	playerInfo.Difficulty = difficulty
	db.UpdateDifficulty(&playerData, difficulty)

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

	lastMove := LastMoveType{
		Row: row + 1,
		Col: col + 1,
		Num: num,
		Type: "MOVE",
	}

	// return the result as json
	return HandleResultLastMove(&playerInfo, &playerData, success, lastMove)
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

	lastMove := LastMoveType{
		Row: row + 1,
		Col: col + 1,
		Num: num,
		Type: "HINT",
	}

	if remove {
		lastMove.Type = "HINT_REMOVE"
	}

	// return the result as json
	return HandleResultLastMove(&playerInfo, &playerData, success, lastMove)
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
		fmt.Println("1. Error quitting game: ", playerErr)
		return "", playerErr
	}

	// get board
	board, boardErr := GetBoard(db, playerInfo)
	if boardErr != nil {
		fmt.Println("2. Error quitting game: ", boardErr)
		return "", boardErr
	}

	playerInfo.FinishBoard()
	playerData.Wins = playerInfo.Wins
	playerData.Losses = playerInfo.Losses
	playerData.PerfectWins = playerInfo.PerfectWins
	playerData.Points = playerInfo.Points
	dbPlayErr := db.UpdatePlayer(&playerInfo)
	if dbPlayErr != nil {
		fmt.Println("3. Error quitting game: ", dbPlayErr)
		return "", dbPlayErr
	}

	board.QuitGame()

	// delete board from database
	dbErr := db.DeleteGame(playerInfo.Id)
	if dbErr != nil {
		fmt.Println("4. Error quitting game: ", dbErr)
		return "", dbErr
	}

	// return the result as json
	return HandleResult(&playerInfo, &playerData, true)
}
