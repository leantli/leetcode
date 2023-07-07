package main

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
// 121. 买卖股票的最佳时机

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp 怎么定义？ prices[i] 表示第 i 天股票的价格
// 定义 dp[i][0] 和 dp[i][1] 表示 第 i 天买或卖股票的利润？
// 但是基于买卖的行为，如果 dp[3][1] 表示第 3 天卖出股票，后面又要如何考虑状态转移呢？
// 因此我们定义 dp[i][0] 和 dp[i][1] 为，第 i 天持有股票和位持有股票的利润
// 持有股票不代表持有当天购入，可能是之前购入的，此时 dp[i][0] = max(dp[i-1][0], -prices[i])
// 为什么 dp[i][0] 不是等于 max(dp[i-1][0], dp[i-1][1]-prices[i]) ? 因为这道题只做一次买卖，我们不需要考虑多次购买
// 那么此时我们要看看哪天购入股票价格更低
// 不持有股票的状态来源，之前也不持有(也可能是之前已经卖了)或者今天卖了 dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0])
// 显然 dp[i]x 只和 dp[i-1]x 相关，因此我们其实再一次优化可以只需要两个变量
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0], dp[0][1] = -prices[0], 0
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0])
	}
	// 1 不持有的状态下收益一定是最高的
	return dp[n-1][1]
}

// // 不用 dp -- 贪心 -- 在当前之前取最低价，当前之后取最高价
// // 用一个 min 记录当前天之前的最低价即可
// func maxProfit(prices []int) int {
// 	minPrice := math.MaxInt
// 	var res int
// 	for _, price := range prices {
// 		if price < minPrice {
// 			minPrice = price
// 		}
// 		res = max(price-minPrice, res)
// 	}
// 	return res
// }
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
