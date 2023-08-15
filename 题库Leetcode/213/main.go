package main

// https://leetcode.cn/problems/house-robber-ii/
// 213. 打家劫舍 II

// 二刷
// 第 i 家偷与不偷，取决于是否要偷上一家，偷上一家钱多还是偷这家钱多
// 定义 dp[i] 为偷前 i 家房屋的最大金额是多少，此时状态转移方程为 dp[i] = max(dp[i-1], dp[i-2]+nums[i])
// 再由于 i 只由 i-1 和 i-2 决定，因此可以用两个局部变量替代 dp[i-1] 和 dp[i-2]
// 而本题的房屋成环，需要考虑首尾的房屋不能同时被偷取，显然，这里首和尾只能取其一甚至极端情况两边都不取
// 这里我们可以分成两次去取 dp，一次剔除掉尾部，一次剔除掉首部，进行 dp 转移，分别获取两次的最大金额
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var a, b int
	for i := range nums[:len(nums)-1] {
		a, b = b, max(b, a+nums[i])
	}
	var c, d int
	for _, num := range nums[1:] {
		c, d = d, max(d, c+num)
	}
	return max(b, d)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// // 相比打家劫舍1，这里多出了房屋成环
// // 但其实主要影响也只有首位两个房屋，因此我们只需要保证，取首时不取尾，取尾时不取首
// // 那么我们可以考虑，仍和之前一样做dp, 但是分成两次，一次预估取首，一次预估取尾
// // 预估取首这次，我们默认绝对不会用尾，则直接取数组[:n-1]；预估取尾这次，默认绝不会用首，取数组[1:]
// // dp[i] = max(dp[i-2]+nums[i], do[i-1])
// func rob(nums []int) int {
// 	n := len(nums)
// 	if n == 1 {
// 		return nums[0]
// 	}
// 	if n == 2 {
// 		return max(nums[0], nums[1])
// 	}
// 	getMax := func(nums []int) int {
// 		ppre, pre := nums[0], max(nums[0], nums[1])
// 		for _, num := range nums[2:] {
// 			ppre, pre = pre, max(pre, ppre+num)
// 		}
// 		return pre
// 	}
// 	return max(getMax(nums[:n-1]), getMax(nums[1:]))
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
