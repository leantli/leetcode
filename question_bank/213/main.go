package main

// https://leetcode.cn/problems/house-robber-ii/
// 213. 打家劫舍 II

// 相比打家劫舍1，这里多出了房屋成环
// 但其实主要影响也只有首位两个房屋，因此我们只需要保证，取首时不取尾，取尾时不取首
// 那么我们可以考虑，仍和之前一样做dp, 但是分成两次，一次预估取首，一次预估取尾
// 预估取首这次，我们默认绝对不会用尾，则直接取数组[:n-1]；预估取尾这次，默认绝不会用首，取数组[1:]
// dp[i] = max(dp[i-2]+nums[i], do[i-1])
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	getMax := func(nums []int) int {
		ppre, pre := nums[0], max(nums[0], nums[1])
		for _, num := range nums[2:] {
			ppre, pre = pre, max(pre, ppre+num)
		}
		return pre
	}
	return max(getMax(nums[:n-1]), getMax(nums[1:]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
