package main

// https://leetcode.cn/problems/number-of-islands/description/
// 200. 岛屿数量

// 首先 双重循环 去看对应位置是否为 陆地1，是则 res 岛屿数量+1，并基于回溯，将该岛屿所有 1 改为 0，避免重复计算
// 遍历完数组则结果出来了
func numIslands(grid [][]byte) int {
	n, m := len(grid), len(grid[0])
	var res int
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if y < 0 || y >= n || x < 0 || x >= m || grid[y][x] == '0' {
			return
		}
		grid[y][x] = '0'
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				res++
				dfs(j, i)
			}
		}
	}
	return res
}
