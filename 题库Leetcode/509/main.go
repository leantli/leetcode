package main

// https://leetcode.cn/problems/fibonacci-number/
// 509. 斐波那契数

// 简单的 dp，已经给出了状态转移方程 dp[n] = dp[n-1] + dp[n-2]
// 并且由于 n 只由 n-1 和 n-2 决定，因此可以用两个局部变量替代
func fib(n int) int {
	if n <= 1 {
		return n
	}
	pre, cur := 0, 1
	// i 从 2 开始，因为 f(0) 和 f(1) 已计算为 pre 和 cur
	for i := 2; i <= n; i++ {
		pre, cur = cur, pre+cur
	}
	return cur
}
