package board

import "fmt"

/* Will allow setting hints */

func (b *Board) SetHint(row, col, num int) bool {
	if num < 1 || num > 9 {
		return false
	}
	if (b.ValidMove(row, col, num)) {
		b.Hints[col][row][num-1] = true
		return true
	}

	return false
}

/* Will allow removing hints */

func (b *Board) RemoveHint(row, col, num int) {
	if num < 1 || num > 9 {
		return
	}
	b.Hints[col][row][num-1] = false
}

/* Will check if a hint is set */

func (b *Board) HasHint(row, col, num int) bool {
	return b.Hints[col][row][num-1]
}

/* Will set a move on the player board, removing the needed hints when doing so if the move is valid */

func (b *Board) SetMove(row, col, num int) bool {
	fmt.Println("SetMove", row, col, num)
	fmt.Println("Result is", b.Board[col][row])
	if (b.Board[col][row] == num) {
		fmt.Println("RIGHT")
		b.PlayerBoard[col][row] = num
		b.RemoveHints(row, col, num)
		return true
	}
	fmt.Println("Board", b.Board);
	fmt.Println("WRONG")

	b.Mistakes++
	return false
}

/* Will remove the hints for a move, both in that space, the 4x4 grid and in the row and col */

func (b *Board) RemoveHints(row, col, num int) {
	for i := 0; i < 9; i++ {
		b.Hints[col][i][num-1] = false
		b.Hints[i][row][num-1] = false
	}

	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.Hints[i+startCol][j+startRow][num-1] = false
		}
	}
}

func (b *Board) BoardDone() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b.PlayerBoard[i][j] == b.Board[i][j] {
				return false
			}
		}
	}

	return true
}

func (b *Board) GetMistakes() int {
	return b.Mistakes
}

func (b *Board) QuitGame() {
	b.EmptyHints()
	b.Mistakes = 0
	b.PlayerBoard = [9][9]int{}
	b.Board = [9][9]int{}
}

