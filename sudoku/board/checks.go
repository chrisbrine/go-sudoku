package board

/* Will check if a number is in a 3x3 box */

func InBox(row, col, num int, board [9][9]int) bool {
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startCol][j+startRow] == num {
				return true
			}
		}
	}
	return false
}

/* Will check if a number is in a row */

func InRow(row, num int, board [9][9]int) bool {
	for i := 0; i < 9; i++ {
		if i != row && board[i][row] == num {
			return true
		}
	}
	return false
}

/* Will check if a number is in a column */

func InCol(col, num int, board [9][9]int) bool {
	for i := 0; i < 9; i++ {
		if i != col && board[col][i] == num {
			return true
		}
	}
	return false
}

/* Will check if a number is in a 4x4 box, row, or column */

func ValidMoveBoard(row, col, num int, board [9][9]int) bool {
	return !InBox(row, col, num, board) && !InRow(row, num, board) && !InCol(col, num, board)
}

func (b *Board) ValidMove(row, col, num int) bool {
	return ValidMoveBoard(row, col, num, b.PlayerBoard)
}

func (b *Board) NumberOfNLeft(number int) int {
	count := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b.PlayerBoard[i][j] == number {
				count++
			}
		}
	}
	return 9 - count
}

func (b *Board) NumbersLeft() [9]int {
	counts := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 9; i++ {
		counts[i] = b.NumberOfNLeft(i + 1)
	}
	return counts
}

func (b *Board) Playing() bool {
	if b.BoardDone() {
		return false
	}

	if b.Mistakes >= 20 {
		return false
	}

	if len(b.Board) == 0 {
		return false
	}

	return true
}

func (b *Board) Win() bool {
	return b.BoardDone() && b.Mistakes < 20
}

func (b *Board) Lose() bool {
	return b.Mistakes >= 20
}
