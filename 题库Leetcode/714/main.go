package main

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
// 714. 买卖股票的最佳时机含手续费

// dp[i][0-1] 定义为 第 i 天 持有/未持有 股票的收益
// 此时加了手续费，卖出操作会降低 fee 价值的收益
// dp[i][0] 第 i 天持有股票收益 来源 --- 第i-1天也持有股票 或 第 i-1 天没持有股票的收益 - 今天买股票开销
// dp[i][1] 第 i 天未持有股票收益 来源 --- 第i-1天也未持有股票 或 第 i-1 天持有股票的收益 + 今天卖股票赚的 - fee
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	// dp[0][1] = 0，因为这一天不会现买现卖，直接从 0 变成 -fee
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}
	return dp[n-1][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
