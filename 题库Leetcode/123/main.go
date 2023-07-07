package main

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
// 123. 买卖股票的最佳时机 III

// 最多只能买两笔，记录最大的两个 dp[i][1]？但显然不行，因为两个 dp[i][1] 可能是同一次交易导致的收益
// 再重新考虑，之前是同一天有两个状态，持有股票和未持有；此时显然可能存在四种状态
// dp[i][0~3] 表示 第 i 天 第一次持有、第一次卖出后未持有、第二次持有、第二次卖出后未持有 的收益
// 第 i 天第一次持有的收益来源 第i-1天第一次持有的收益，或者今天刚买入的收益 dp[i][0] = max(dp[i-1][0], -prices[i])
// 第 i 天第一次未持有的收益来源 第i-1天第一次未持有的收益，或者第i-1天第一次持有后今天卖出 dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
// 第 i 天第二次持有的收益来源 第i-1天第二次持有的收益，或者第一次未持有后的收益减去今天刚买入价格的收益 dp[i][2] = max(dp[i-1][2], dp[i-1][1]-prices[i])
// 第 i 天第二次未持有的收益来源 第i-1天第二次未持有的收益，或者第i-1天第二次持有后今天卖出 dp[i][3] = max(dp[i-1][3], dp[i-1][2]+prices[i])
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 4)
	}
	// 注意第 0 天第二次卖出也还是 -prices[0]，即第一天买入了再卖出然后又买入了
	dp[0][0], dp[0][1], dp[0][2] = -prices[0], 0, -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]-prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]+prices[i])
	}
	return dp[n-1][3]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
