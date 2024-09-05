package board

// create a new game

func Create(difficulty int) *Board {
	board := &Board{}

	// Initialize the board
	board.InitBoard()
	board.SetupBoard(difficulty)
	board.EmptyHints()

	return board
}

func CreateFromJson(data string) *Board {
	return fromJson(data)
}
