package game

import (
	"encoding/json"

	"github.com/chrisbrine/go-sudoku/sudoku/sql"
)

func GetLeaderboard(db *sql.DBData) (string, error) {
	// get leaderboard
	dbLeaderboard, dbErr := db.GetLeaderboard()
	if dbErr != nil {
		return "", dbErr
	}

	// convert to leaderboard
	leaderboard := []Leader{}
	for _, dbPlayer := range dbLeaderboard {
		leaderboard = append(leaderboard, Leader{
			Username: dbPlayer.Username,
			Name: dbPlayer.Name,
			Wins: dbPlayer.Wins,
			Losses: dbPlayer.Losses,
			Points: dbPlayer.Points,
			PerfectWins: dbPlayer.PerfectWins,
		})
	}

	// convert leaderboard to json string
	jsonLeaderboard, jsonErr := json.Marshal(leaderboard)
	if jsonErr != nil {
		return "", jsonErr
	}

	// return leaderboard
	return string(jsonLeaderboard), nil
}