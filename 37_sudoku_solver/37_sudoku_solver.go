package sudokusolver

func solveSudoku(board [][]byte) {
	backtracking(board)
}

func backtracking(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			//判断此位置是否适合填数字
			if board[i][j] != '.' {
				continue
			}
			//尝试填1-9
			for k := '1'; k <= '9'; k++ {
				if isvalid(i, j, byte(k), board) { //如果满足要求就填
					board[i][j] = byte(k)
					if backtracking(board) {
						return true
					}

					board[i][j] = '.'
				}
			}
			return false
		}
	}
	return true
}

//判断填入数字是否满足要求
func isvalid(row, col int, k byte, board [][]byte) bool {
	for i := 0; i < 9; i++ { //行
		if board[row][i] == k {
			return false
		}
	}

	for i := 0; i < 9; i++ { //列
		if board[i][col] == k {
			return false
		}
	}

	//方格
	startrow := (row / 3) * 3
	startcol := (col / 3) * 3
	for i := startrow; i < startrow+3; i++ {
		for j := startcol; j < startcol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}

	return true
}
