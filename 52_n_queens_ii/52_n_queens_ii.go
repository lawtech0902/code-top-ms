package nqueensii

func totalNQueens(n int) int {
	res := 0
	chessboard := make([][]rune, n)
	for i := range chessboard {
		item := make([]rune, n)
		for j := range item {
			item[j] = '.'
		}

		chessboard[i] = item
	}

	backTracking(0, chessboard, &res)
	return res
}

func backTracking(row int, chessboard [][]rune, res *int) {
	if row == len(chessboard) {
		*res++
		return
	}

	for i := 0; i < len(chessboard); i++ {
		if isValid(row, i, chessboard) {
			chessboard[row][i] = 'Q'
			backTracking(row+1, chessboard, res)
			chessboard[row][i] = '.'
		}
	}
}

func isValid(row, col int, chessboard [][]rune) bool {
	// 检查同列
	for i := 0; i < row; i++ {
		if chessboard[i][col] == 'Q' {
			return false
		}
	}

	// 检查45度角斜线
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == 'Q' {
			return false
		}
	}

	// 检查135度角斜线
	for i, j := row-1, col+1; i >= 0 && j < len(chessboard); i, j = i-1, j+1 {
		if chessboard[i][j] == 'Q' {
			return false
		}
	}

	return true
}
