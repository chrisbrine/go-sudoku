package board

/* Will allow setting hints */

func (b *Board) SetHint(row, col, num int) bool {
	if (b.ValidMove(row, col, num)) {
		b.hints[row][col][num-1] = true
		return true
	}

	return false
}

/* Will allow removing hints */

func (b *Board) RemoveHint(row, col, num int) {
	b.hints[row][col][num-1] = false
}

/* Will check if a hint is set */

func (b *Board) HasHint(row, col, num int) bool {
	return b.hints[row][col][num-1]
}

/* Will set a move on the player board, removing the needed hints when doing so if the move is valid */

func (b *Board) SetMove(row, col, num int) bool {
	if (b.ValidMove(row, col, num)) {
		b.playerBoard[row][col] = num
		b.RemoveHints(row, col, num)
		return true
	}

	b.mistakes++
	return false
}

/* Will remove the hints for a move, both in that space, the 4x4 grid and in the row and col */

func (b *Board) RemoveHints(row, col, num int) {
	for i := 0; i < 9; i++ {
		b.hints[row][i][num-1] = false
		b.hints[i][col][num-1] = false
	}

	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.hints[i+startRow][j+startCol][num-1] = false
		}
	}
}

func (b *Board) BoardDone() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b.playerBoard[i][j] == b.board[i][j] {
				return false
			}
		}
	}

	return true
}

func (b *Board) GetMistakes() int {
	return b.mistakes
}

