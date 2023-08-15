package main

// https://leetcode.cn/problems/integer-break/
// 343. 整数拆分

// 和剑指offer的剪绳子是一样的
func integerBreak(n int) int {
	// n 可以拆成 1 ~ n-1，定义 dp[i] 为整数 i 拆分后乘积最大为多少
	// 此时 dp[n] = max{dp[1]*n-1, dp[2]*n-2, ... dp[n-n+1]*(n-1), 1 * (n-1), 2 * (n-2), ... (n-n+1)*(n-1)}
	dp := make([]int, n+1)
	// 这里 i 从 2 开始，但是我们需要保证 dp[2] = 1
	// 所以要初始化 dp[1] = 1，单纯由状态转移方程逆推而出，保证方程正常运作，本身没有实际意义，因为 i = 1 时无法拆分
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(max(dp[i], j*dp[i-j]), j*(i-j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
