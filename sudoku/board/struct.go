/* Create the game board struct */

package board

import "encoding/json"

type Board struct {
	Board [9][9]int `json:"Board"`
	Hints [9][9][9]bool `json:"Hints"`
	PlayerBoard [9][9]int `json:"PlayerBoard"`
	Mistakes int `json:"Mistakes"`
}

func (b *Board) toJson() string {
	data, _ := json.Marshal(b)
	return string(data)
}

func fromJson(data string) *Board {
	board := Board{}
	json.Unmarshal([]byte(data), &board)
	return &board
}
