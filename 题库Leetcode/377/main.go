package main

// https://leetcode.cn/problems/combination-sum-iv/
// 377. 组合总和 Ⅳ

// 每个数可以无限取，找到总和为 target，并且不是组合而是排列
// 首先肯定是完全背包 dp[j] 为 从 0-i 选取物品，放入容量为 j 的背包中，有多少种排列
// 则 dp[j] += dp[j-nums[i]]，滚动数组优化后降维
// 由于是求排列而不是组合，将物品放入内循环，j 背包容量在外循环
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	// 初始化，保证后续状态转移正常
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
