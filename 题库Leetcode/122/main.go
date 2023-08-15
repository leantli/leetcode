package main

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
// 122. 买卖股票的最佳时机 II

// 二刷
// 可以在同一天购入和售出，并且可以多次买卖股票
// 显然使用 dp 解决, 定义 dp[i][j] 表示第 i 天持有或未持有股票的收益
// dp[i][0] 表示第 i 天未持有股票的收益 dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])，之前也没持有，或前一天持有，今天卖掉
// dp[i][1] 表示第 i 天持有股票的收益 dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])，之前也持有，或前一天未持有，今天买入
// 由于 i 只由 i-1 的持有和未持有推出，我们无需维持整个 dp 数组，只需要维持两个局部变量即可
// func maxProfit(prices []int) int {
//     uncontrol, control := 0, -prices[0]
//     for i := 1; i < len(prices); i++ {
//         uncontrol, control = max(uncontrol, control + prices[i]), max(control, uncontrol-prices[i])
//     }
//     return uncontrol
// }
// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

// 看了之前的解答，发现还有一个很牛的思路，因为可以当天买卖并且可以多次买入卖出，显然，只要今天的比昨天贵，就能买入卖出，一定有的赚
func maxProfit(prices []int) int {
	var res int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

// 和 121 相比较，在任意时间段仍然只能持有一股股票，但多出了可以反复购买的情况考虑

// // 只针对这道题的贪心---只要今天的股价比昨天的高，就买卖！因为可以同一天买或者同一卖
// // 所以永远只会赚！
// func maxProfit(prices []int) int {
// 	var res int
// 	for i := 1; i < len(prices); i++ {
// 		if prices[i] > prices[i-1] {
// 			res += prices[i] - prices[i-1]
// 		}
// 	}
// 	return res
// }

// // 仍然是 dp[i][0] 和 dp[i][1]，表示第 i 天持有或未持有股票的收益
// // 此时持有股票 dp[i][0] 的来源，昨天也持有股票，昨天没持有但今天买了 dp[i][0] = max(dp[i-1][0], dp[i][1] - prices[i])
// // 此时未持有股票 dp[i][1] 的来源，昨天也没持有，昨天持有但今天卖了 dp[i][0] = max(dp[i-1][0], dp[i-1][0] + prics[i]])
// func maxProfit(prices []int) int {
// 	n := len(prices)
// 	dp := make([][]int, n)
// 	for i := range dp {
// 		dp[i] = make([]int, 2)
// 	}
// 	dp[0][0], dp[0][1] = -prices[0], 0
// 	for i := 1; i < n; i++ {
// 		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
// 		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
// 	}
// 	return dp[n-1][1]
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
