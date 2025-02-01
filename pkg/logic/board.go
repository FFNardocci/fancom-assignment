package logic

func InitBoard(size int) [][]string {
	board := make([][]string, size)

	//Assign helper color values on initiating board
	for i := 0; i < size; i++ {
		//Allocate inner slice
		board[i] = make([]string, size)
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				board[i][j] = "b"
			} else {
				board[i][j] = "w"
			}
		}
	}
	return board
}
