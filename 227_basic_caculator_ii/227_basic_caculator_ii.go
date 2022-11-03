package basiccaculatorii

import "unicode"

// + - * /
// 栈
func calculate(s string) int {
	stack := make([]int, 0)
	num := 0
	sign := '+'

	for i, r := range s {
		if unicode.IsDigit(r) {
			num = num*10 + int(r-'0')
			if i != len(s)-1 {
				continue
			}
		}

		// 跳过空格
		if r == ' ' && i != len(s)-1 {
			continue
		}

		switch sign {
		case '+':
			stack = append(stack, num)
		case '-':
			stack = append(stack, -num)
		case '*':
			newNum := stack[len(stack)-1] * num
			stack[len(stack)-1] = newNum
		case '/':
			newNum := stack[len(stack)-1] / num
			stack[len(stack)-1] = newNum
		}

		num = 0
		sign = r
	}

	res := 0
	for _, item := range stack {
		res += item
	}

	return res
}
