package main

// https://leetcode.cn/problems/longest-continuous-increasing-subsequence/
// 674. 最长连续递增序列

// 最长 递增的子序列
// 不定长滑动窗口，窗口内性质为递增
// 此时我们需要注意，每次右扩时，将该数与上一个数比较(连续递增，看题中举例)
// 不满足时，l 并不是右移一位，而是直接移动至当前 r 所处的位置
// 初始化时要注意，窗口内得有初始数
func findLengthOfLCIS(nums []int) int {
	var l int
	maxLen := 1
	for r := 1; r < len(nums); r++ {
		if nums[r] <= nums[r-1] {
			l = r
		}
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 数组未排序，找到 最长+递增 连续子序列 返回其长度
// 这种连续的子序列，显然是可以考虑基于 滑动窗口 去度量去最长长度
// func findLengthOfLCIS(nums []int) int {
// 	// 设置 l,r, ans 的起始下标
// 	var l int
// 	ans := 1
// 	// 我们要寻找连续递增，必定需要将 nums[r]及nums[r-1]进行比较，因此直接设 r 起始位为 1
// 	for r := 1; r < len(nums); r++ {
// 		// 不满足条件时，窗口直接销毁，从新开始，因此 l=r，以此保证条件满足递增
// 		if nums[r] <= nums[r-1] {
// 			l = r
// 		}
// 		// 满足条件就正常计算长度
// 		ans = max(ans, r-l+1)
// 	}
// 	return ans
// }
