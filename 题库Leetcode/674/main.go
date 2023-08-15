package main

// https://leetcode.cn/problems/longest-continuous-increasing-subsequence/
// 674. 最长连续递增序列

// 二刷
// 这里要求的子序列需要连续，算是子数组？连续递增子数组
// 直接判断连续的最长的递增子数组的长度即可
func findLengthOfLCIS(nums []int) int {
	res, cnt := 1, 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			cnt++
			if cnt > res {
				res = cnt
			}
		} else {
			cnt = 1
		}
	}
	return res
}

// 二刷
// 数组未排序，要找最长的连续递增的子序列，注意此处有连续，比正常子序列多了连续的限制
// 算是 LIS 的简化版本，因为这里连续，因此我们可以利用不定长滑动窗口解决
// 窗口性质---递增
// func findLengthOfLCIS(nums []int) int {
// 	var l, r int
// 	res := 1
// 	for r < len(nums)-1 {
// 		// 维护窗口性质
// 		if nums[r+1] <= nums[r] {
// 			l = r + 1
// 		}
// 		r++
// 		if res < r-l+1 {
// 			res = r - l + 1
// 		}
// 	}
// 	return res
// }

// 考虑在训练dp专题，这题显然也是可以用dp来解，算是T300的弱化版，T300需要考虑全部子序列(包括不连续的情况)
// 这里只需要考虑连续的情况
// 显然我们还是定义 dp[i] --> 以nums[i]结尾的最长连续递增序列的长度
// 初始化：全部dp[i]置为1，因为所有数字单独都能成为连续递增子序列
// 状态转移：
// dp[i] = dp[i-1]+1, nums[i]>nums[i-1]时
//
//	func findLengthOfLCIS(nums []int) int {
//		res := 1
//		dp := make([]int, len(nums))
//		for i := range nums {
//			dp[i] = 1
//		}
//		for i := 1; i < len(nums); i++ {
//			if nums[i] > nums[i-1] {
//				dp[i] = dp[i-1] + 1
//			}
//			if dp[i] > res {
//				res = dp[i]
//			}
//		}
//		return res
//	}
//
// // 并且我们注意，这个dp，其实只需要两个状态, i 和 i-1，因此我们完全无需维护一个dp数组
// func findLengthOfLCIS(nums []int) int {
// 	res, pre := 1, 1
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] > nums[i-1] {
// 			pre += 1
// 		} else {
// 			pre = 1
// 		}
// 		if pre > res {
// 			res = pre
// 		}
// 	}
// 	return res
// }

// // 最长 递增的子序列
// // 不定长滑动窗口，窗口内性质为递增
// // 此时我们需要注意，每次右扩时，将该数与上一个数比较(连续递增，看题中举例)
// // 不满足时，l 并不是右移一位，而是直接移动至当前 r 所处的位置
// // 初始化时要注意，窗口内得有初始数
// func findLengthOfLCIS(nums []int) int {
// 	var l int
// 	maxLen := 1
// 	for r := 1; r < len(nums); r++ {
// 		if nums[r] <= nums[r-1] {
// 			l = r
// 		}
// 		maxLen = max(maxLen, r-l+1)
// 	}
// 	return maxLen
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

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
