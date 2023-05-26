package main

// https://leetcode.cn/problems/the-masseuse-lcci/
// 面试题 17.16. 按摩师

// 相邻不能取，求最大，类似 198 打家劫舍
// dp[i] 表示 前 i 个预约，最长达到的预约时长
// 如果取第 i 个预约，则第i-1个预约不能取，dp[i] = dp[i-2]+nums[i]
// 如果不去第 i 个预约，则 dp[i] = dp[i-1]
// 因此 dp[i] = max(dp[i-1], dp[i-2]+nums[i])
// 关注到 i 只与 i-1、i-2 相关，我们可以用两个局部遍历替代整个 dp 数组的声明
func massage(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	ppre, pre := nums[0], max(nums[0], nums[1])
	for _, num := range nums[2:] {
		ppre, pre = pre, max(pre, ppre+num)
	}
	return pre
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
