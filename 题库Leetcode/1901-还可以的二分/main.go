package main

// https://leetcode.cn/problems/find-a-peak-element-ii/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=bv79h7h
// 1901. 寻找峰值 II

// 再重新考虑一下，其实峰值应该是一定存在的？
// 二维数组里面最大的必然就是峰值，但是这样可能需要 m*n 的时间复杂度
// 我们显然拒绝暴力遍历，之前对每行求峰值，实际上并不是每行的最大值，存在遗漏
// 因此每行找最大值，组成一个数组后再二分找峰值，但是又不能遍历完全部找每行的最大值，否则还不如全遍历找全部数的最大值
// 因此我们只需要找到每行最大值数组前面的峰值即可
func findPeakGrid(mat [][]int) []int {
	m := len(mat)
	var lastMax, lastIndex int
	for i := 0; i < m; i++ {
		curMax, curIndex := getMaxAndIndex(mat[i])
		if curMax < lastMax {
			return []int{i - 1, lastIndex}
		}
		lastMax, lastIndex = curMax, curIndex
	}
	return []int{m - 1, lastIndex}
}

func getMaxAndIndex(nums []int) (max, index int) {
	for i, num := range nums {
		if num > max {
			max = num
			index = i
		}
	}
	return
}

// // 暴力不考虑，题目要求 时间复杂度为 O(m log(n)) 或 O(n log(m))
// // 找到每行的峰值，再与其上下比较？这样也算是 nlogm 或者 mlogn
// // 思路错误，同一行可能存在多个峰值，但我们一行只会得到一个峰值，
// // 这意味着我们可能会错失同一行其他可能能够成为峰值的行峰值
// func findPeakGrid(mat [][]int) []int {
// 	m, n := len(mat), len(mat[0])
// 	for i := 0; i < m; i++ {
// 		l, r := 0, n-1
// 		for l < r {
// 			mid := l + (r-l)/2
// 			if mat[i][mid] < mat[i][mid+1] {
// 				l = mid + 1
// 			} else {
// 				r = mid
// 			}
// 		}
// 		if judgeTop(mat, l, i) {
// 			return []int{i, l}
// 		}
// 	}
// 	return []int{-1, -1}
// }
// func judgeTop(mat [][]int, x, y int) bool {
// 	top := mat[y][x]
// 	// 左
// 	if x-1 >= 0 && top < mat[y][x-1] {
// 		return false
// 	}
// 	// 右
// 	if x+1 < len(mat[0]) && top < mat[y][x+1] {
// 		return false
// 	}
// 	// 上
// 	if y-1 >= 0 && top < mat[y-1][x] {
// 		return false
// 	}
// 	// 下
// 	if y+1 < len(mat) && top < mat[y+1][x] {
// 		return false
// 	}
// 	return true
// }
