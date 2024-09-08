/* Generate a full finished board at random */

package board

import (
	"fmt"
	"math/rand"
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
		if b.PlayerBoard[row][col] == 0 {
			continue
		}

		// Remove the number
		b.PlayerBoard[row][col] = 0
		numberToRemove--
	}
}

func (b *Board) SetupBoard(difficulty int) {
	// This runs after the board is created. Go through the board and, based on the difficulty level, set the necessary
	// numbers in playerBoard to 0 and the rest to the number in board

	numberDifficulties := []int{30, 30, 40, 50}
	useDifficulty := numberDifficulties[0]

	// Set the player board to the board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b.PlayerBoard[i][j] = b.Board[i][j]
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

func (b *Board) Swap3X3Blocks(r1 int, r2 int, block [9][9]int) [9][9]int {
	for i := 0; i < 3; i++ {
		block = b.SwapRows(r1 * 3 + i, r2 * 3 + i, block)
	}
	return block
}

func (b *Board) SwapCols(c1 int, c2 int, board[9][9]int) [9][9]int {
	for i := 0; i < 9; i++ {
		col := board[i][c1]
		board[i][c1] = board[i][c2]
		board[i][c2] = col
	}
	return board
}

func (b *Board) SwapRows(r1 int, r2 int, board [9][9]int) [9][9] int {
	row := board[r1]
	board[r1] = board[r2]
	board[r2] = row
	return board
}

func (b *Board) SwapNumbers(n1 int, n2 int, board [9][9]int) [9][9]int {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == n1 {
				board[i][j] = n2
			} else if board[i][j] == n2 {
				board[i][j] = n1
			}
		}
	}
	return board
}

func (b *Board) Shuffle3X3Blocks(board [9][9]int) [9][9]int {
	for i := 0; i < 3; i++ {
		ranNum := rand.Intn(3)
		board = b.Swap3X3Blocks(i, ranNum, board)
	}
	return board
}

func (b *Board) ShuffleCols(board [9][9]int) [9][9]int {
	var blockNumber int
	for i := 0; i < 3; i++ {
		ranNum := rand.Intn(3)
		blockNumber = i / 3
		board = b.SwapCols(i, blockNumber * 3 + ranNum, board)
	}
	return board
}

func (b *Board) ShuffleRows(board [9][9]int) [9][9]int {
	var blockNumber int
	for i := 0; i < 3; i++ {
		ranNum := rand.Intn(3)
		blockNumber = i / 3
		board = b.SwapRows(i, blockNumber * 3 + ranNum, board)
	}
	return board
}

func (b *Board) ShuffleNumbers(board [9][9]int) [9][9]int {
	for i := 1; i <= 9; i++ {
		ranNum := rand.Intn(9) +1
		board = b.SwapNumbers(i, ranNum, board)
	}
	return board
}

func (b *Board) ShuffleAll(board [9][9]int) [9][9]int {
	board = b.ShuffleNumbers(board)
	board = b.ShuffleRows(board)
	board = b.ShuffleCols(board)
	board = b.Shuffle3X3Blocks(board)
	return board
}

func (b *Board) InitBoard() {
	startingBoard := [9][9]int{
		{1,2,3,  4,5,6,  7,8,9},
		{4,5,6,  7,8,9,  1,2,3},
		{7,8,9,  1,2,3,  4,5,6},

		{2,3,1,  5,6,4,  8,9,7},
		{5,6,4,  8,9,7,  2,3,1},
		{8,9,7,  2,3,1,  5,6,4},

		{3,1,2,  6,4,5,  9,7,8},
		{6,4,5,  9,7,8,  3,1,2},
		{9,7,8,  3,1,2,  6,4,5},
	};
	b.Board = b.ShuffleAll(startingBoard)
	fmt.Println("Board", b.Board)
}

func (b *Board) EmptyHints() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k++ {
				b.Hints[i][j][k] = false
			}
		}
	}
}
