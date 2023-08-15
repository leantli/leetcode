package main

// https://leetcode.cn/problems/coin-change/description/
// 322. 零钱兑换

// 背包问题 dp[i][j] 表示从 0~i 个物品选取，放入背包中，价值最高是多少
// dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
// 再通过滚动数组优化降维,dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
// 而这里不是选取价值最高，而是选取达到 amount 的最少硬币个数
// 则定义 dp[j] 为从 0~i 个物品选取，放入背包中，达到背包重量 j 时，最少需要的硬币个数是多少
// 则 dp[j] = min(dp[j], dp[j-weight[i]]+1)，并且由于这是完全背包，每个值可以反复取，因此可以无需倒序遍历
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	max := amount + 1 // coin最小为 1，那使用的硬币数量最多就是 amount，不可能是 amount+1，此处仅做最大值初始化
	for i := range dp {
		dp[i] = max
	}
	// 同时 dp[0] 需要初始化为 0，否则状态转移无法正常进行
	dp[0] = 0
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coin]+1)
		}
	}
	if dp[amount] == max {
		return -1
	}
	return dp[amount]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
