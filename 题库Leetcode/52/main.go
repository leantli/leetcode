package main

// https://leetcode.cn/problems/n-queens-ii/
// 52. N 皇后 II

// dfs 回溯
// 返回所有 n 皇后的放置方案
// 我们可以使用 used[i] 表示第 i 列是否有皇后已经占据了
// 然后 for i := 0; i < n; i++ 逐行去放置皇后，这样能够保证同一行不会出现两个皇后
// 此时我们解决了行和列上的皇后重复问题，再写个函数去判断皇后不会在斜线上重复即可
func totalNQueens(n int) int {
	used := make([]bool, n)
	grid := make([][]bool, n)
	for i := range grid {
		grid[i] = make([]bool, n)
	}
	var res int
	// y 表示当前放置皇后的行数
	var dfs func(y int)
	dfs = func(y int) {
		if y == n {
			res++
			return
		}
		for j := 0; j < n; j++ {
			if !used[j] && canPlace(j, y, grid) {
				used[j] = true
				grid[y][j] = true
				dfs(y + 1)
				used[j] = false
				grid[y][j] = false
			}
		}
	}
	dfs(0)
	return res
}

// 判断皇后能否放置，判断是否斜线上已经有皇后了
// 由于我们是自上往下放置皇后的，因此斜线上的皇后的判断，我们只需要判断上面行数的即可
func canPlace(x, y int, grid [][]bool) bool {
	n := len(grid)
	for i := 0; i <= y; i++ {
		// 只要 (x,y) 坐标左上斜线和右上斜线存在皇后，则直接返回 false 表示不能放置皇后
		if x-i >= 0 && grid[y-i][x-i] {
			return false
		}
		if x+i < n && grid[y-i][x+i] {
			return false
		}
	}
	return true
}
