package main

// https://leetcode.cn/problems/longest-continuous-increasing-subsequence/
// 674. 最长连续递增序列

// 数组未排序，找到 最长+递增 连续子序列 返回其长度
// 这种连续的子序列，显然是可以考虑基于 滑动窗口 去度量去最长长度
func findLengthOfLCIS(nums []int) int {
	// 设置 l,r, ans 的起始下标
	var l int
	ans := 1
	// 我们要寻找连续递增，必定需要将 nums[r]及nums[r-1]进行比较，因此直接设 r 起始位为 1
	for r := 1; r < len(nums); r++ {
		// 不满足条件时，窗口直接销毁，从新开始，因此 l=r，以此保证条件满足递增
		if nums[r] <= nums[r-1] {
			l = r
		}
		// 满足条件就正常计算长度
		ans = max(ans, r-l+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
