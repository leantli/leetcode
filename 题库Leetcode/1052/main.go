package main

// https://leetcode.cn/problems/grumpy-bookstore-owner/
// 1052. 爱生气的书店老板

// minutes 显然是一个限定的长度
// 这显然也可以基于滑动窗口去解决这道题
// 我们只需要判断 minutes 窗口中，grumpy = 1 的情况下，反转为 0 后，会多出多少顾客满意
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	var more, base, maxMore int
	for i, num := range customers[:minutes] {
		if grumpy[i] == 1 {
			more += num
		} else {
			base += num
		}
	}
	maxMore = more
	for r := minutes; r < len(customers); r++ {
		l := r - minutes
		if grumpy[l] == 1 {
			more -= customers[l]
		}
		if grumpy[r] == 1 {
			more += customers[r]
		} else {
			base += customers[r]
		}
		maxMore = max(maxMore, more)
	}
	return base + maxMore
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
