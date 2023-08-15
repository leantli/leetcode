package main

// https://leetcode.cn/problems/coin-change-ii/description/
// 518. 零钱兑换 II

// 完全背包问题，硬币为不同的物品，可无限取，背包最大负重为 amount
// dp[i][j] += dp[i-1][j-coins[i]] + 1?
// 再根据滚动数组优化降维，dp[j] += dp[j-coins[i]]
// 并且由于是完全背包，可无限取，因此无需倒序遍历，正序遍历即可
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	// dp[0] 需要初始化为 1，保证每个 dp[coins[i]] 在第一次遍历时，能够加上 1，保证自身 dp[coins[i]] 有一种组合
	dp[0] = 1
	// 由于题目要求求硬币组合，排序在其中无所谓，所以我们将硬币放在外部循环
	// 如果题目要求求硬币排列，则将背包放在外部，内循环才是硬币，保证每个amount都依次选取过各个 coin
	for _, coin := range coins {
		// 由于是可重复取 coin，所以这里正序遍历，即在尝试重复取 coin
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}
	return dp[amount]
}
