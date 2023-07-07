package main

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
// 188. 买卖股票的最佳时机 IV

// // dp[i][j] = 第 i 天，第 j/2+1 次持有/未持有股票时的收益
// // 能够买卖 k 次，意味着有 2*k 次持有和未持有的状态
// func maxProfit(k int, prices []int) int {
// 	n := len(prices)
// 	dp := make([][]int, n)
// 	for i := range dp {
// 		dp[i] = make([]int, 2*k)
// 	}
// 	// 第一天只有有持有股票，收益一定为 -prices[0]，不管当天交易买入卖出多少次
// 	for i := 0; i < 2*k; i++ {
// 		if i&1 == 0 {
// 			dp[0][i] = -prices[0]
// 		}
// 	}
// 	for i := 1; i < n; i++ {
// 		dp[i][0] = max(dp[i-1][0], -prices[i])
// 		for j := 1; j < 2*k; j++ {
// 			temp := prices[i]
// 			// 如果是买入，第 k 次持有，就将价格改为负数，当 j 为偶数时，都是持有状态
// 			if j&1 == 0 {
// 				temp = -temp
// 			}
// 			dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+temp)
// 		}
// 	}
// 	return dp[n-1][2*k-1]
// }

// dp[i][j] = 第 i 天，第 j/2+1 次持有/未持有股票时的收益
// 能够买卖 k 次，意味着有 2*k 次持有和未持有的状态
// 再简单优化一下空间
func maxProfit(k int, prices []int) int {
	n := len(prices)
	dp := make([]int, 2*k)
	// 第一天只有有持有股票，收益一定为 -prices[0]，不管当天交易买入卖出多少次
	for i := 0; i < 2*k; i++ {
		if i&1 == 0 {
			dp[i] = -prices[0]
		}
	}
	for i := 1; i < n; i++ {
		for j := 2*k - 1; j > 0; j-- {
			temp := prices[i]
			// 如果是买入，第 k 次持有，就将价格改为负数，当 j 为偶数时，都是持有状态
			if j&1 == 0 {
				temp = -temp
			}
			dp[j] = max(dp[j], dp[j-1]+temp)
		}
		dp[0] = max(dp[0], -prices[i])
	}
	return dp[2*k-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
