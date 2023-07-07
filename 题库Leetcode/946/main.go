package main

// https://leetcode.cn/problems/validate-stack-sequences/
// 946. 验证栈序列

// 模拟？我们可以用 指针 a 遍历 popped
// 当 a 指向的变量，不在 pushed 中存在时，按 pushed 顺序正常将对应的数入栈
// 入栈时遇到 a 指向的变量时，a 向右移动并不断 pushed 出栈
// 当 pushed 的序列都用完了，再看看 a 走没走到最后，走到最后就说明栈序列没问题
func validateStackSequences(pushed []int, popped []int) bool {
	temp := make([]int, 0)
	for _, num := range pushed {
		temp = append(temp, num)
		for len(popped) > 0 && len(temp) > 0 && temp[len(temp)-1] == popped[0] {
			popped = popped[1:]
			temp = temp[:len(temp)-1]
		}
	}
	return len(popped) == 0
}
