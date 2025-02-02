package logic

// Leaving this here mostly for signifying a possible future extension of the application
func InitBoard(size int) [][]string {
	board := make([][]string, size)
	for i := 0; i < size; i++ {
		board[i] = make([]string, size)
	}

	return board
}
