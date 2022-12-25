package searcha2dmatrix

import "sort"

// 两次二分, 不知道为啥过不了
func searchMatrix(matrix [][]int, target int) bool {
	rowIndex := binarySearchFirstColumn(matrix, target)
	if rowIndex < 0 {
		return false
	}

	return binarySearchRow(matrix[rowIndex], target)
}

func binarySearchFirstColumn(matrix [][]int, target int) int {
	low, high := -1, len(matrix)-1
	for low < high {
		mid := low + (high-low+1)/2
		if matrix[mid][0] <= target {
			low = mid
		} else {
			high = mid - 1
		}
	}

	return low
}

func binarySearchRow(row []int, target int) bool {
	low, high := 0, len(row)-1
	for low <= high {
		mid := low + (high-low)/2
		if row[mid] == target {
			return true
		} else if row[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}

// 两次二分 sort
func searchMatrix1(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool {
		return matrix[i][0] > target
	}) - 1
	if row < 0 {
		return false
	}

	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}

// 一次二分
func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i := sort.Search(m*n, func(i int) bool {
		return matrix[i/n][i%n] >= target
	})

	return i < m*n && matrix[i/n][i%n] == target
}
