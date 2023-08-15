package main

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// 309. 最佳买卖股票时机含冷冻期

// 二刷
// dp[i][j] 表示第 i 天持有/未持有/陷入冷冻期的股票的收益
// 第 i 天持有，dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i]) 第 i 天持有，可能是前一天也持有，或者前一天是冷冻期，今天可以买入
// 第 i 天未持有,dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]) 第 i 天未持有，可能是前一天也未持有，也可能是前一天持有，今天卖出
// 第 i 天陷入冷冻期,dp[i][2] = dp[i-1][1] 第 i 天陷入冷冻期，一定是第 i-1 天卖出变成未持有
// 由于 i 只由 i-1 决定，显然每一天三个状态的收益可以由三个局部变量替代
func maxProfit(prices []int) int {
	var control, uncontrol, freeze int
	control = -prices[0]
	for i := 1; i < len(prices); i++ {
		control, uncontrol, freeze = max(control, freeze-prices[i]), max(uncontrol, control+prices[i]), uncontrol
	}
	return uncontrol
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// // dp[i][0~1] 定义为 第 i 天持有/未持有股票的收益
// // 第 i 天持有股票的收益 来源---第 i-1 天持有股票的收益 或 第i-2天未持有股票的收益 - 今天购入的开销
// // 注意，这里不是从 第 i-1 天未持有股票 - 今天购入开销，因为第 i-1 天可能存在冷冻期
// // 第 i 天未持有股票的收益 来源-- 第 i-1 天未持有股票收益 或 第i-1天持有股票 + 今天卖出的收益
// func maxProfit(prices []int) int {
// 	n := len(prices)
// 	if n <= 1 {
// 		return 0
// 	}
// 	dp := make([][]int, n)
// 	for i := range dp {
// 		dp[i] = make([]int, 2)
// 	}
// 	dp[0][0] = -prices[0]
// 	dp[1][0] = max(-prices[0], -prices[1])
// 	dp[1][1] = max(dp[0][1], dp[0][0]+prices[1])
// 	for i := 2; i < n; i++ {
// 		dp[i][0] = max(dp[i-1][0], dp[i-2][1]-prices[i])
// 		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
// 	}
// 	return dp[n-1][1]
// }

// // 当然我们也可以多给一个冷冻期的状态
// // dp[i][0~2] 定义为 第 i 天持有、未持有、冷冻期(冷冻期之后)的收益
// // 第 i 天持有股票的收益 来源---第 i-1 天持有股票的收益 或 第i-1天冷冻期的收益 - 今天购入的开销
// // 第 i 天未持有股票的收益 来源-- 第 i-1 天未持有股票收益 或 第i-1天持有股票 + 今天卖出的收益
// // 第 i 天冷冻期收益 来源-- 第i-1天未持有股票的收益(卖出)
// func maxProfit(prices []int) int {
// 	n := len(prices)
// 	if n <= 1 {
// 		return 0
// 	}
// 	dp := make([][]int, n)
// 	for i := range dp {
// 		dp[i] = make([]int, 3)
// 	}
// 	dp[0][0] = -prices[0]
// 	for i := 1; i < n; i++ {
// 		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
// 		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
// 		dp[i][2] = dp[i-1][1]
// 	}
// 	return dp[n-1][1]
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
