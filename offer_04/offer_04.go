package offer04

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	row, col := 0, cols-1
	for row < rows && col >= 0 {
		num := matrix[row][col]
		if num == target {
			return true
		} else if num > target {
			col--
		} else {
			row++
		}
	}

	return false
}
