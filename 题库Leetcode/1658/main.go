package main

// https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/
// 1658. 将 x 减到 0 的最小操作数

// 这种做法也不错，和之前的思路差不多，只是之前我的思路是计算窗口内的值
// 这种做法就是直接计算窗口外的值，更减少了一步转换
func minOperations(nums []int, x int) int {
	n, total := len(nums), 0
	for _, num := range nums {
		total += num
	}
	if total < x {
		return -1
	}
	right := 0
	ans := n + 1
	lSum, rSum := 0, total
	for left := -1; left < n; left++ {
		if left >= 0 {
			lSum += nums[left]
		}
		for right < n && lSum+rSum > x {
			rSum -= nums[right]
			right++
		}

		if lSum+rSum == x {
			if left+1+n-right < ans {
				ans = left + 1 + n - right
			}
		}
	}
	if ans > n {
		return -1
	}
	return ans
}

// // 考虑成两个区域？不定长滑窗
// // 首先我们需要计算整个数组和，为 sum
// // 中间的区域即为滑窗，为不移除的窗口，窗口内的值满足条件为 sum - x，此时窗口外(移除的值)和为 x
// // 当窗口内的值 和为 sum-x,此时计算窗口的最大长度，窗口长度最长，即需要移除的操作最少
// // 当和 小于 sum-x，r右移扩大窗口，大于 sum-x时，l 右移缩小窗口
// func minOperations(nums []int, x int) int {
// 	var sum int
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	if sum < x {
// 		return -1
// 	}
// 	if sum == x {
// 		return len(nums)
// 	}
// 	target := sum - x
// 	var l, r, sumWindow, maxLen int
// 	for r < len(nums) {
// 		sumWindow += nums[r]
// 		for sumWindow > target {
// 			sumWindow -= nums[l]
// 			l++
// 		}
// 		r++
// 		if sumWindow == target {
// 			maxLen = max(maxLen, r-l)
// 		}
// 	}
// 	if maxLen == 0 {
// 		return -1
// 	}
// 	return len(nums) - maxLen
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
