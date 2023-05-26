package main

// https://leetcode.cn/problems/house-robber/
// 198. 打家劫舍

// 每个房屋相邻不能被偷，如何定义状态呢？
// dp[i] 定义为什么？i 房屋有偷和不偷两个状态
// 假设我要偷第 i 家, dp[i] = nums[i] + dp[i-2]
// 假设我不偷第 i 家， dp[i] = dp[i-1]
// 那么前 i 家的最高偷窃金额为 dp[i] = max(dp[i-1], dp[i-2]+nums[i])
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	// 并且，我们发现，dp[i]只由i-1和i-2两个决定，因此可以使用两个局部变量替代dp数组的使用
	ppre, pre := nums[0], max(nums[0], nums[1])
	for _, num := range nums[2:] {
		ppre, pre = pre, max(pre, ppre+num)
	}
	return pre
}

// f[i] 偷第i家，则 f[i] = g[i-1]+nums[i]
// g[i] 不偷第i家，则第 i-1 家可偷可不偷，则 g[i] = max(f[i-1], g[i-1])
// 则 dp[i] = max(f[i],g[i])
// 参考第 337 题的想法，想到的状态转移方程，但是为什么官解不用这种思路而是 dp[i] = max(dp[i-1], dp[i-2]+nums[i])
// 因为这里是一维数组情况，操作 i, i-1, i-2 其实比较方便，因此直接一套即可
// 而 337 中 i 为父节点，i-1 为两个子节点， i-2 可能是四个孙子节点，一来数量上不太好直接 max，二来不像一维数组取值方便
// 因而通过定义 f[i] 和 g[i]，分别为第i家偷/不偷进行操作和做转移方程，会更简单
// func rob(nums []int) int {
// 	var f, g int
// 	for _, num := range nums {
// 		f, g = g+num, max(f, g)
// 	}
// 	return max(f, g)
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
