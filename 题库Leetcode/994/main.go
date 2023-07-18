package main

// https://leetcode.cn/problems/rotting-oranges/
// 994. 腐烂的橘子

type position struct {
	x, y int
}

func orangesRotting(grid [][]int) int {
	var fresh, day int
	queue := make([]position, 0)
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				queue = append(queue, position{x: j, y: i})
			}
		}
	}
	if fresh == 0 {
		return 0
	}
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	for len(queue) > 0 {
		day++
		times := len(queue)
		for i := 0; i < times; i++ {
			for j := 0; j < 4; j++ {
				x, y := queue[i].x+dx[j], queue[i].y+dy[j]
				if x >= 0 && x < m && y >= 0 && y < n && grid[y][x] == 1 {
					grid[y][x] = 2
					fresh--
					queue = append(queue, position{x, y})
				}
			}
		}
		queue = queue[times:]
	}
	if fresh == 0 {
		return day - 1
	}
	return -1
}
