package basiccaculator

import "unicode"

// + - ()
// 栈
func calculate(s string) int {
	stack := make([]int, 0)
	res := 0
	sign := 1

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			// 暂存左边数字
			num := 0
			for ; i < len(s); i++ {
				if !unicode.IsDigit(rune(s[i])) {
					break
				}

				num = num*10 + int(s[i]-'0')
			}

			res += sign * num
			i--
		} else {
			switch s[i] {
			case '+':
				sign = 1
			case '-':
				sign = -1
			case '(':
				// 遇左括号，结果及符号入栈并重置
				stack = append(stack, res, sign)
				res, sign = 0, 1
			case ')':
				// 遇有括号，取栈中暂存结果及符号计算新的结果
				sign = stack[len(stack)-1]
				num := stack[len(stack)-2]
				res = num + sign*res
				stack = stack[:len(stack)-2]
			}
		}
	}

	return res
}
