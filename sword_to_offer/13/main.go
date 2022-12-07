package main

// https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 13. 机器人的运动范围

var res int

func movingCount(m int, n int, k int) int {
	// 存储 x,y 坐标点是否已判断过
	mm := make([][]bool, m, m)
	for i := 0; i < m; i++ {
		mm[i] = make([]bool, n, n)
	}
	res = 0
	helper(mm, m, n, 0, 0, k)
	return res
}

func helper(mm [][]bool, m, n, x, y, k int) {
	// 先判断边界，已经判断过的也直接返回
	if x < 0 || x >= n || y < 0 || y >= m || mm[y][x] == true {
		return
	}
	mm[y][x] = true
	if !isVaild(x, y, k) {
		return
	}
	res++
	helper(mm, m, n, x+1, y, k)
	helper(mm, m, n, x, y+1, k)
	helper(mm, m, n, x, y-1, k)
	helper(mm, m, n, x-1, y, k)
}

// 判断该格子是否能进入
func isVaild(x, y, k int) bool {
	sum := 0
	for x != 0 {
		sum += x % 10
		x /= 10
	}
	for y != 0 {
		sum += y % 10
		y /= 10
	}
	return sum <= k
}
