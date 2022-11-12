package nqueens

import "strings"

func solveNQueens(n int) [][]string {
	var res [][]string
	chessboard := make([][]string, n)
	for i := range chessboard {
		item := make([]string, n)
		for j := range item {
			item[j] = "."
		}

		chessboard[i] = item
	}

	backTracking(0, chessboard, &res)
	return res
}

func backTracking(row int, chessboard [][]string, res *[][]string) {
	if row == len(chessboard) {
		temp := make([]string, len(chessboard))
		for i, rowStrArr := range chessboard {
			temp[i] = strings.Join(rowStrArr, "")
		}

		*res = append(*res, temp)
	}

	for i := 0; i < len(chessboard); i++ {
		if isValid(row, i, chessboard) {
			chessboard[row][i] = "Q"
			backTracking(row+1, chessboard, res)
			chessboard[row][i] = "."
		}
	}
}

func isValid(row, col int, chessboard [][]string) bool {
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < len(chessboard); i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}

	return true
}
