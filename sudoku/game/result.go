package game

import (
	"encoding/json"

	"github.com/chrisbrine/go-sudoku/sudoku/board"
	"github.com/chrisbrine/go-sudoku/sudoku/player"
	"github.com/chrisbrine/go-sudoku/sudoku/sql"
)

/*
Get player data and board

takes commands with DB to get player and board data along with any commands such as:
- new game
- get / refresh current game and player data
- set number in row and col
- set hint in row and col

These will ALWAYS return a json data of the current player and board data and it will always be the same data
*/

func PlayerDataToResult(playerData *player.Player, DBplayer *sql.DBPlayer, success bool, lastMove LastMoveType) GameResult {
	gameBoard := playerData.GetGame()

	var playerBoard [9][9]int
	var hints [9][9][9]bool
	var mistakes int
	var numbersLeft [9]int
	var playing bool
	var inGame bool
	gameStatus := 0

	if gameBoard != nil {
		playerBoard = gameBoard.PlayerBoard
		hints = gameBoard.Hints
		mistakes = gameBoard.Mistakes
		numbersLeft = gameBoard.NumbersLeft()
		playing = gameBoard.Playing()
		inGame = true
		if !gameBoard.Playing() {
			if (gameBoard.Win()) {
				gameStatus = 1
			} else {
				gameStatus = -1
			}
		}
	} else {
		playerBoard = [9][9]int{}
		hints = [9][9][9]bool{}
		mistakes = 0
		numbersLeft = [9]int{}
		playing = false
		inGame = false
	}

	result := &GameResult{
		Board: playerBoard,
		// PlayerBoard: gameBoard.PlayerBoard,
		Hints: hints,
		Mistakes: mistakes,
		NumbersLeft: numbersLeft,
		Playing: playing,
		InGame: inGame,
		Points: playerData.GetPoints(),
		Wins: playerData.GetWins(),
		Losses: playerData.GetLosses(),
		GameStatus: gameStatus,
		LastMove: lastMove,
		PerfectWins: playerData.GetPerfectWins(),
		Difficulty: playerData.GetDifficulty(),
		Username: DBplayer.GetUsername(),
		Name: DBplayer.GetName(),
		Success: success,
	}

	return *result
}

func ResultToJson(result *GameResult) (string, error) {
	// change to json
	jsonData, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func HandleResultLastMove(playerData *player.Player, DBplayer *sql.DBPlayer, success bool, lastMove LastMoveType) (string, error) {
	result := PlayerDataToResult(playerData, DBplayer, success, lastMove)

	jsonData, err := ResultToJson(&result)
	if err != nil {
		return "", err
	}

	return jsonData, nil
}

func HandleResult(playerData *player.Player, DBplayer *sql.DBPlayer, success bool) (string, error) {
	return HandleResultLastMove(playerData, DBplayer, success, LastMoveType{})
}

func GetBoard(db *sql.DBData, player player.Player) (*board.Board, error) {
	game, boardErr := db.GetGame(player.Id)
	if boardErr != nil {
		return &board.Board{}, boardErr
	}

	player.SetGame(game)

	return game, nil
}
