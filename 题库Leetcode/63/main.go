package main

// https://leetcode.cn/problems/unique-paths-ii/
// 63. 不同路径 II

// 这道题显然可以通过 dfs 解决，但题目只要得到不同的路径有多少，那么显然我们可以采用更好的方法
// 定义 dp[i][j] 为坐标 i,j 位置有多少不同的路径到达, dp[m-1][n-1] 显然就是最终答案
// 而状态转移方法即是 dp[i][j] = dp[i-1][j] + dp[i][j], if grid[i][j] != 1，只要位置没有障碍，就正常计算，否则为 0
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	// 初始化，没遇到障碍时都为 1，遇到障碍则为 0
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 0 {
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
