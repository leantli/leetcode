package main

// https://leetcode.cn/problems/count-negative-numbers-in-a-sorted-matrix/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 1351. 统计有序矩阵中的负数

// // 注意该矩阵，行和列都是非递增
// // 可以直接二分，找到每行第一个负数即可，时间复杂度为 O(mlogn)，测试一下
// // 要找到每行的负数，左边是非负数，右边是负数，采用万金油模板比较舒服
// func countNegatives(grid [][]int) int {
// 	var ans int
// 	n := len(grid[0])
// 	for _, raw := range grid {
// 		l, r := -1, n
// 		for l+1 != r {
// 			mid := l + (r-l)/2
// 			if raw[mid] >= 0 {
// 				l = mid
// 			} else {
// 				r = mid
// 			}
// 		}
// 		// 此时出循环后，l 为每行最后一个非负数，r 为该行第一个负数
// 		// 当然，也可能没有负数，导致 r 仍为 n
// 		if r == n {
// 			continue
// 		}
// 		// 否则 n-r 就是该行有的负数数量
// 		ans += n - r
// 	}
// 	return ans
// }

// 逐行二分可以解决问题，但显然不满足题目进阶要求的时间复杂度 O(m+n)
// 再额外关注一下，逐行是递减的，逐列也是递减的
// 我们使用二分也只用到了 行递减，没用到 列递减 的特性
// 这里其实要注意观察到，由于 行递减 以及 列递减 的特性
// 导致越往下，数组出现负数的下标只会越来越往左，即便不左移，其下标至少是上一行负数出现的下标
// 因此我们可以从右上角开始向左遍历，遇到非负数时下移，再继续左移至该行的非负数，如此反复
// 此时时间复杂度为 O(n+m)，x 和 y 的移动最多只会是 n+m 次
func countNegatives(grid [][]int) int {
	var ans int
	n := len(grid[0])
	x, y := n-1, 0
	for y < len(grid) {
		for x >= 0 && grid[y][x] < 0 {
			x--
		}
		// 此时 x 会停在当前行的最后一个非负数处
		ans += n - x - 1
		y++
	}
	return ans
}
