package nqueens

import (
	"strings"
)

func solveNQueens(n int) [][]string {
	var res [][]string
	chessboard := make([][]string, n)
	for i := range chessboard {
		chessboard[i] = make([]string, n)
	}

	for i := range chessboard {
		for j := range chessboard {
			chessboard[i][j] = "."
		}
	}

	backTracking(n, 0, chessboard, &res)
	return res
}

func backTracking(n, row int, chessboard [][]string, res *[][]string) {
	if row == n {
		temp := make([]string, n)
		for i, rowStr := range chessboard {
			temp[i] = strings.Join(rowStr, "")
		}

		*res = append(*res, temp)
	}

	for i := 0; i < n; i++ {
		if isValid(n, row, i, chessboard) {
			chessboard[row][i] = "Q"
			backTracking(n, row+1, chessboard, res)
			chessboard[row][i] = "."
		}
	}
}

func isValid(n, row, col int, chessboard [][]string) bool {
	// 检查同列
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}

	// 检查45度角斜线
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}

	// 检查135度角斜线
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}

	return true
}
