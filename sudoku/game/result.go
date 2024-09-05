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

func PlayerDataToResult(playerData *player.Player, DBplayer *sql.DBPlayer, success bool) GameResult {
	gameBoard := playerData.GetGame()

	result := &GameResult{
		Board: gameBoard.Board,
		PlayerBoard: gameBoard.PlayerBoard,
		Hints: gameBoard.Hints,
		Mistakes: gameBoard.Mistakes,
		NumbersLeft: gameBoard.NumbersLeft(),
		Playing: gameBoard.Playing(),
		Points: playerData.GetPoints(),
		Wins: playerData.GetWins(),
		Losses: playerData.GetLosses(),
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

func HandleResult(playerData *player.Player, DBplayer *sql.DBPlayer, success bool) (string, error) {
	result := PlayerDataToResult(playerData, DBplayer, success)

	jsonData, err := ResultToJson(&result)
	if err != nil {
		return "", err
	}

	return jsonData, nil
}

func GetBoard(db *sql.DBData, player player.Player) (*board.Board, error) {
	game, boardErr := db.GetGame(player.Id)
	if boardErr != nil {
		return &board.Board{}, boardErr
	}

	player.SetGame(game)

	return game, nil
}
