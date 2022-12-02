package main

// https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 47. 礼物的最大价值

// dp 或者 dfs 都行，但是这道题无需说明有多少路径，并且只要最大价值，最佳选 dp + 原地修改
// dp[y][x] = max { dp[y-1][x], dp[y][x-1] }
func maxValue(grid [][]int) int {
	// 初始化边界
	xMax, yMax := len(grid[0]), len(grid)
	for i := 1; i < xMax; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < yMax; i++ {
		grid[i][0] += grid[i-1][0]
	}
	// 状态转移计算
	for y := 1; y < yMax; y++ {
		for x := 1; x < xMax; x++ {
			grid[y][x] = Max(grid[y-1][x], grid[y][x-1]) + grid[y][x]
		}
	}
	return grid[yMax-1][xMax-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
