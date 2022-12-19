package main

// https://leetcode.cn/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 31. 栈的压入、弹出序列

// 应该就是模拟，
// 每添加一个数进栈，就配合 popped 数组判断能不能出栈
// 直到不能再出栈，再重复上面的步骤
func validateStackSequences(pushed []int, popped []int) bool {
	mock := make([]int, 0)
	for _, num := range pushed {
		mock = append(mock, num)
		for len(popped) > 0 && len(mock) > 0 && mock[len(mock)-1] == popped[0] {
			mock = mock[:len(mock)-1]
			popped = popped[1:]
		}
	}
	return len(popped) == 0
}
