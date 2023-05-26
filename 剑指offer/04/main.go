package main

// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 4. 二维数组中的查找

// // 常规做法，看到左到右非递减，上到下非递减，是有其规律的，从二维数组右上角看，非常像个二叉搜索树
// // 可以从行最大，列最小的位置开始对角线排查; 列最大，行最小 O(N+M)
// func findNumberIn2DArray(matrix [][]int, target int) bool {
// 	if len(matrix) == 0 {
// 		return false
// 	}
// 	x, y := len(matrix[0])-1, 0
// 	for x >= 0 && y < len(matrix) {
// 		if matrix[y][x] == target {
// 			return true
// 		}
// 		if matrix[y][x] > target {
// 			x--
// 			continue
// 		}
// 		if matrix[y][x] < target {
// 			y++
// 		}
// 	}
// 	return false
// }

// 暴力不多说 O(NM)
