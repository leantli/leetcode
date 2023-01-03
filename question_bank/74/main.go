package main

// https://leetcode.cn/problems/search-a-2d-matrix/
// 74. 搜索二维矩阵

// 首先我们知道每行是升序，每列更是必定是升序(每行第一个整数必定大于前一行最后一个整数)
// 因此我们可以从最右上角开始，若比 target 小，则列增，若比 target 大，则行减
// 如此一来，每次递增递减都可以排除一半的选择
// func searchMatrix(matrix [][]int, target int) bool {
// 	if len(matrix) == 0 || len(matrix[0]) == 0 {
// 		return false
// 	}
// 	y, x := 0, len(matrix[0])-1
// 	for x >= 0 && y < len(matrix) {
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

// 直接二分，将二维数组模拟成一维进行二分
// 如何将二维数组变成一维数组操作呢？
// 二维数组的总长为 n * m (len(martrix) * len(matrix[0]))
// 二维数组拉成一维数组，第 k 个数，位于 k/n 的数组上的 k%n 下标上
// 比如，下面是一个二维数组
//
//	    0 1 2
//		3 4 5
//		6 7 8
//
// 6 位于 6/3=2 数组上的第 6%3=0 下标
func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	l, r := -1, n*m
	for l+1 != r {
		mid := l + (r-l)/2
		if matrix[mid/m][mid%m] < target {
			l = mid
		} else {
			r = mid
		}
	}
	// 后处理一下
	if r == n*m {
		return false
	}
	if matrix[r/m][r%m] != target {
		return false
	}
	return true
}
