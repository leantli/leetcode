package main

import "strings"

// https://leetcode.cn/problems/decode-string/
// 394. 字符串解码

func decodeString(s string) string {
	stack := make([]byte, 0)
	NumberStack := make([]int, 0)
	for i := 0; i < len(s); i++ {
		// 对数字单独处理，处理完后 i 移动到 [ 下标继续正常操作
		if s[i] >= '0' && s[i] <= '9' {
			var temp int
			for s[i] >= '0' && s[i] <= '9' {
				temp = temp*10 + int(s[i]-'0')
				i++
			}
			NumberStack = append(NumberStack, temp)
		}
		// 不是右括号就正常进栈
		if s[i] != ']' {
			stack = append(stack, s[i])
			continue
		}
		// 遇到右括号，不会对右括号进栈，反而要将栈中，出栈顺序遇到的第一个 [ 后的字符都串起来
		var temp string
		for stack[len(stack)-1] != '[' {
			temp = string(stack[len(stack)-1]) + temp
			stack = stack[:len(stack)-1]
		}
		stack = stack[:len(stack)-1] // 排除掉 '['
		temp = strings.Repeat(temp, NumberStack[len(NumberStack)-1])
		NumberStack = NumberStack[:len(NumberStack)-1]
		stack = append(stack, temp...)
	}
	return string(stack)
}
