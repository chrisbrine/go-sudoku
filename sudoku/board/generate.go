/* Generate a full finished board at random */

package board

import (
	"math/rand"
	"time"
)

func (b *Board) removeNumbers(numberToRemove int) {
	// This will remove the necessary number of numbers from the player board to create the puzzle
	// it will also make sure it doesn't keep removing numbers that are in the same row and column so a number is only removed once
	// it sets to 0 when removed

	for numberToRemove > 0 {
		// Get a random row and column
		row := rand.Intn(9)
		col := rand.Intn(9)

		// Check if the cell is already empty
		if b.playerBoard[row][col] == 0 {
			continue
		}

		// Remove the number
		b.playerBoard[row][col] = 0
		numberToRemove--
	}
}

func (b *Board) SetupBoard(difficulty int) {
	// This runs after the board is created. Go through the board and, based on the difficulty level, set the necessary
	// numbers in playerBoard to 0 and the rest to the number in board

	numberDifficulties := []int{40, 40, 50, 60}
	useDifficulty := numberDifficulties[0]

	// Set the player board to the board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b.playerBoard[i][j] = b.board[i][j]
		}
	}

	// Set the numbers in the player board to 0 based on the difficulty level

	// check if difficulty is less than 0 or greater than the length, if so set it to 0
	if difficulty > 0 && difficulty < len(numberDifficulties) {
		useDifficulty = numberDifficulties[difficulty]
	}

	b.removeNumbers(useDifficulty)

}

// Generate a full finished board at random

// randomValidNumber returns a random valid number for the cell at the given row and column
func (b *Board) randomizeNumber(row, col int) int {
	// Create a slice of all possible numbers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Shuffle the slice
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })

	// Go through each number and check if it is valid
	for _, number := range numbers {
		if b.ValidMove(row, col, number) {
			return number
		}
	}

	// If no valid number was found, return 0
	return 0
}

func (b *Board) InitBoard() {
	// Go through each cell in the board and fill it with a random valid number
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b.board[i][j] = b.randomizeNumber(i, j)
		}
	}
}

func (b *Board) EmptyHints() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k++ {
				b.hints[i][j][k] = false
			}
		}
	}
}
