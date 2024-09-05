/* Create the game board struct */

package board

import "encoding/json"

type Board struct {
	Board [9][9]int `json:"Board"`
	Hints [9][9][9]bool `json:"Hints"`
	PlayerBoard [9][9]int `json:"PlayerBoard"`
	Mistakes int `json:"Mistakes"`
}

func (b *Board) ToJson() (string, error) {
	data, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func fromJson(data string) (*Board, error) {
	board := Board{}
	err := json.Unmarshal([]byte(data), &board)
	if err != nil {
		return nil, err
	}
	return &board, nil
}
