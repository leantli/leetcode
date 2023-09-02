package main

// https://leetcode.cn/problems/max-area-of-island/description/
// 695. 岛屿的最大面积

// 双重 for 循环去判断是否遇到陆地，是则回溯去计算整个陆地有多大，并且将陆地值置为0，避免重复计算陆地面积
func maxAreaOfIsland(grid [][]int) int {
	var res int
	n, m := len(grid), len(grid[0])
	var dfs func(x, y int)
	var area int
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || grid[y][x] == 0 {
			return
		}
		area++
		grid[y][x] = 0
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}
	for i := range grid {
		for j := range grid[i] {
			// 如果是陆地，则开始计算陆地面积并将陆地置为 0
			if grid[i][j] == 1 {
				area = 0
				dfs(j, i)
				if area > res {
					res = area
				}
			}
		}
	}
	return res
}
