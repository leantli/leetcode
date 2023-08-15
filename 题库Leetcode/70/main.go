package main

// https://leetcode.cn/problems/climbing-stairs/
// 70. 爬楼梯

// 本质和 509 斐波那契数列一致
func climbStairs(n int) int {
	// 定义 dp[i] 为到达第 i 阶有多少种不同的方法
	// 由于每次都只能爬 1 或 2 个台阶，显然 dp[i] = dp[i-1] + dp[i-2]
	// 而初始化 dp[1] = 1, dp[2] = 2，接着自底向上，得到 dp[n]
	// 并且由于 dp[i] 只由 i-1 和 i-2 决定，可由局部变量替代
	if n <= 2 {
		return n
	}
	pre, cur := 1, 2
	// i 从 3 开始，因为 1 和 2 已经初始化
	for i := 3; i <= n; i++ {
		pre, cur = cur, pre+cur
	}
	return cur
}
