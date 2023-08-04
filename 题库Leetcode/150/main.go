package main

import "strconv"

// https://leetcode.cn/problems/evaluate-reverse-polish-notation/
// 150. 逆波兰表达式求值

// 一个简单的小优化，使用 strconv 转数字是否成功判断 token 是否为数字
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		// 当 token 能正常转成数字，则直接入栈，否则进行计算
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
			continue
		}
		// 否则进行计算，将计算结果也加入栈中，因为后续仍要使用到
		// 这里注意，a 是第二栈顶数，b 是第一栈顶数，不要弄反了
		a, b := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		switch token {
		case "+":
			stack = append(stack, a+b)
		case "-":
			stack = append(stack, a-b)
		case "*":
			stack = append(stack, a*b)
		case "/":
			stack = append(stack, a/b)
		}
	}
	// 最终栈中剩下的数即为结果
	return stack[0]
}

// func evalRPN(tokens []string) int {
// 	stack := make([]int, 0)
// 	m := map[string]struct{}{
// 		"+": {},
// 		"-": {},
// 		"*": {},
// 		"/": {},
// 	}
// 	for _, token := range tokens {
// 		// 是数字就直接入栈
// 		if _, ok := m[token]; !ok {
// 			num, _ := strconv.Atoi(token)
// 			stack = append(stack, num)
// 			continue
// 		}
// 		// 这里注意，a 是第二栈顶数，b 是第一栈顶数，不要弄反了
// 		a, b := stack[len(stack)-2], stack[len(stack)-1]
// 		stack = stack[:len(stack)-2]
// 		// 否则进行计算，将计算结果也加入栈中，因为后续仍要使用到
// 		switch token {
// 		case "+":
// 			stack = append(stack, a+b)
// 		case "-":
// 			stack = append(stack, a-b)
// 		case "*":
// 			stack = append(stack, a*b)
// 		case "/":
// 			stack = append(stack, a/b)
// 		}
// 	}
// 	// 最终栈中剩下的数即为结果
// 	return stack[0]
// }
