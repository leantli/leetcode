package main

// https://leetcode.cn/problems/unique-paths/
// 62. 不同路径

func uniquePaths(m int, n int) int {
	// 可以根据 dfs 去枚举遍历有多少不同的路径，但显然有更好的方法
	// 机器人到某一个格子，只能是从格子上方或左方移动而来，那显然，到达最右下角，由右下角的左侧和上侧到达
	// 定义 dp[i][j] 为到达 i,j 下标，有多少不同的路径
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]，
	// 初始化则是将 i=0 以及 j=0 这两边的边界值初始化一遍即可
	// n 是 列，m 是行
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
