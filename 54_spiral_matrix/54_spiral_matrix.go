package spiralmatrix

import "reflect"

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	for len(matrix) != 0 {
		res = append(res, matrix[0]...)
		matrix = matrix[1:]
		if len(matrix) == 0 {
			break
		}

		matrix = turn(matrix)
	}

	return res
}

func turn(matrix [][]int) [][]int {
	rowSize, colSize := len(matrix), len(matrix[0])
	var newMatrix [][]int

	for i := 0; i < colSize; i++ {
		var subArr []int
		for j := 0; j < rowSize; j++ {
			subArr = append(subArr, matrix[j][i])
		}

		newMatrix = append(newMatrix, subArr)
	}

	reverseAnySlice(newMatrix)
	return newMatrix
}

func reverseAnySlice(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
