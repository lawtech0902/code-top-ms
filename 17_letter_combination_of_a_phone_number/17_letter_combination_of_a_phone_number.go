package lettercombinationofaphonenumber

func letterCombinations(digits string) []string {
	size := len(digits)
	if size == 0 || size > 4 {
		return nil
	}

	digitsArr := [10]string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}

	res := make([]string, 0)
	backTracking(digitsArr, digits, "", 0, &res)
	return res
}

func backTracking(digitsArr [10]string, digits, tempString string, index int, res *[]string) {
	if len(tempString) == len(digits) {
		*res = append(*res, tempString)
		return
	}

	tempIndex := digits[index] - '0'
	letters := digitsArr[tempIndex]
	for i := 0; i < len(letters); i++ {
		tempString += string(letters[i])
		backTracking(digitsArr, digits, tempString, index+1, res)
		tempString = tempString[:len(tempString)-1]
	}
}
