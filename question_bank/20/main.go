package main

// https://leetcode.cn/problems/valid-parentheses/
// 20. 有效的括号

func isValid(s string) bool {
	// 栈中放左括号
	stack := make([]byte, 0)
	m := make(map[byte]byte)
	m[')'], m[']'], m['}'] = '(', '[', '{'
	for i := range s {
		// 当 s[i] 不是右括号时，存入栈中
		if _, ok := m[s[i]]; !ok {
			stack = append(stack, s[i])
			continue
		}
		// 遇到右括号时，判断右括号对应的左括号是否是栈顶括号
		// 栈中没有左括号或类型不一致，直接返回错误
		if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] {
			return false
		}
		// 否则弹栈，表示匹配成功
		stack = stack[:len(stack)-1]
	}
	// 当栈中还有左括号没匹配时，说明这不是有效括号
	if len(stack) > 0 {
		return false
	}
	return true
}
