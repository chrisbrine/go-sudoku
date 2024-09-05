package board

/* Will check if a number is in a 3x3 box */

func (b *Board) InBox(row, col, num int) bool {
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.Board[i+startRow][j+startCol] == num {
				return true
			}
		}
	}
	return false
}

/* Will check if a number is in a row */

func (b *Board) InRow(row, num int) bool {
	for i := 0; i < 8; i++ {
		if i != row && b.Board[row][i] == num {
			return true
		}
	}
	return false
}

/* Will check if a number is in a column */

func (b *Board) InCol(col, num int) bool {
	for i := 0; i < 8; i++ {
		if i != col && b.Board[i][col] == num {
			return true
		}
	}
	return false
}

/* Will check if a number is in a 4x4 box, row, or column */

func (b *Board) ValidMove(row, col, num int) bool {
	return !b.InBox(row, col, num) && !b.InRow(row, num) && !b.InCol(col, num)
}
