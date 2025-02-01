package logic

func InitBoard(size int) [][]string {
	board := make([][]string, size)

	//Assign helper color values on initiating board
	for i := 0; i < size; i++ {
		board[i] = make([]string, size)
	}

	return board
}
