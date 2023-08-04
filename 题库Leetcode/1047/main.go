package main

// https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/
// 1047. 删除字符串中的所有相邻重复项

// 常规思路就是不断地去遍历字符串，遇到相邻重复项就删除，直到某次遍历没遇到相邻重复项则返回结果
// 这样需要遍历多次，我们可以借用栈去解决这个问题，只需要遍历一次即可
func removeDuplicates(s string) string {
	stack := make([]byte, 0)
	for i := range s {
		// 当栈顶的字符和当前字符相同，说明为相邻重复项，弹出这个元素，并且忽略当前元素入栈
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return string(stack)
}
