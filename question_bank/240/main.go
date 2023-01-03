package main

// https://leetcode.cn/problems/search-a-2d-matrix-ii/
// 240. 搜索二维矩阵 II

// 逐行二分，找到 <= target 的下标
// 当 l 指向的值不等于 target，则列上递增一次并再一次二分找到 <= target
// 时间复杂度是 n * log(m)
func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	for i := 0; i < n; i++ {
		l, r := -1, m
		for l+1 != r {
			mid := l + (r-l)/2
			if matrix[i][mid] <= target {
				l = mid
			} else {
				r = mid
			}
		}
		if l == -1 {
			continue
		}
		if matrix[i][l] == target {
			return true
		}
	}
	return false
}

// // 还有一种就还是从右上角开始，大则列增，小则行减
// // 时间复杂度为 n+m
// func searchMatrix(matrix [][]int, target int) bool {
// 	n, m := len(matrix), len(matrix[0])
// 	x, y := m-1, 0
// 	for x >= 0 && y < n {
// 		if matrix[y][x] == target {
// 			return true
// 		} else if matrix[y][x] > target {
// 			x--
// 		} else {
// 			y++
// 		}
// 	}
// 	return false
// }
