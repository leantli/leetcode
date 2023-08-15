package main

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
// 121. 买卖股票的最佳时机

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 二刷 使用 dp
// 定义 dp[i][j], i 表示第 i 天获利情况，j 只取 0 或 1，0 表示未持有，1 表示持有
// dp[i][0] = dp[i-1][0] || dp[i-1][1] + prices[i] // 未持有的话，要么前一天也未持有，要么前一天持有，今天卖出
// dp[i][1] = dp[i-1][1] || - prices[i] // 持有的话，要么前一天也持有，要么前一天没持有，今天买入
// 由于 i 只由 i-1 的持有和未持有状态决定，因此可以将其简化成两个局部变量，无需使用 dp 数组
// 使用 uncontrol, control 表示未持有和持有
// 则 uncontrol, control = max(uncontrol, control + prices[i]), max(control, -prices[i]),
func maxProfit(prices []int) int {
	var uncontrol, control int
	control = -prices[0]
	for i := 1; i < len(prices); i++ {
		uncontrol, control = max(uncontrol, control+prices[i]), max(control, -prices[i])
	}
	return uncontrol
}

// // 二刷，不使用 dp
// // 目标：买入一次，卖出一次，什么情况下获利最多
// // 我们可以假设每一天都尝试卖出动作，记录这一天之前的最低价格，计算这天卖出是否获利最多即可
// func maxProfit(prices []int) int {
//     lowest := prices[0]
//     var res int
//     for i := 1; i < len(prices); i++ {
//         if prices[i] - lowest > res {
//             res = prices[i] - lowest
//         }
//         if prices[i] < lowest {
//             lowest = prices[i]
//         }
//     }
//     return res
// }

// dp 怎么定义？ prices[i] 表示第 i 天股票的价格
// 定义 dp[i][0] 和 dp[i][1] 表示 第 i 天买或卖股票的利润？
// 但是基于买卖的行为，如果 dp[3][1] 表示第 3 天卖出股票，后面又要如何考虑状态转移呢？
// 因此我们定义 dp[i][0] 和 dp[i][1] 为，第 i 天持有股票和位持有股票的利润
// 持有股票不代表持有当天购入，可能是之前购入的，此时 dp[i][0] = max(dp[i-1][0], -prices[i])
// 为什么 dp[i][0] 不是等于 max(dp[i-1][0], dp[i-1][1]-prices[i]) ? 因为这道题只做一次买卖，我们不需要考虑多次购买
// 那么此时我们要看看哪天购入股票价格更低
// 不持有股票的状态来源，之前也不持有(也可能是之前已经卖了)或者今天卖了 dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0])
// func maxProfit(prices []int) int {
// 	n := len(prices)
// 	dp := make([][]int, n)
// 	for i := range dp {
// 		dp[i] = make([]int, 2)
// 	}
// 	dp[0][0], dp[0][1] = -prices[0], 0
// 	for i := 1; i < n; i++ {
// 		dp[i][0] = max(dp[i-1][0], -prices[i])
// 		dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0])
// 	}
// 	// 1 不持有的状态下收益一定是最高的
// 	return dp[n-1][1]
// }

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
