package main

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
// 188. 买卖股票的最佳时机 IV

// 二刷
// 最多可以完成 k 笔交易，此时显然，每一天都有 k * 2 种交易状态，并且其实我们可以注意到
// dp[i][0] 表示 第 i 天第一次持有的收益, dp[i][0] = max(dp[i-1][0], -prices[i]) 前一天也持有或前一天还没买入，今天买入
// dp[i][1] 表示 第 i 天第一次未持有的收益, dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]) 前一天未持有或前一天持有，今天卖出
// dp[i][2] 表示 第 i 天第二次持有的收益, dp[i][2] = max(dp[i-1][2], dp[i-1][1]-prices[i]) 前一天第二次持有或前一天第一次未持有后今天买入
// dp[i][3] 表示 第 i 天第二次未持有的收益, dp[i][3] = max(dp[i-1][3], dp[i-1][2]+prices[i])
// dp[i][4] 表示 第 i 天第三次持有的收益, dp[i][4] = max(dp[i-1][4], dp[i-1][3]-prices[i])
// 观察到，i 只由 i-1 的状态去决定，因此我们只需要维护 k*2 长度的一维数组即可
// dp[j] = max(dp[j], dp[j-1] +/- prices[i]) // j [0~2*k-1], 0 2 4 等偶数，都是减去 prices[i]，奇数都是加上 prices[i]
// j 更大的依赖于更小的，因此我们显然每一天的 dp[j] 的数据更新，需要从后往前遍历，否则从前往后的话，小的状态已经被覆盖了
func maxProfit(k int, prices []int) int {
	dp := make([]int, 2*k)
	// 初始化，第 0 天，所有的买入收益，都是 -prices[i]
	for i := 0; i < len(dp); i += 2 {
		dp[i] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := 2*k - 1; j > 0; j-- {
			if (j & 1) == 1 {
				dp[j] = max(dp[j], dp[j-1]+prices[i])
			} else {
				dp[j] = max(dp[j], dp[j-1]-prices[i])
			}
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

// // dp[i][j] = 第 i 天，第 j/2+1 次持有/未持有股票时的收益
// // 能够买卖 k 次，意味着有 2*k 次持有和未持有的状态
// // 再简单优化一下空间
// func maxProfit(k int, prices []int) int {
// 	n := len(prices)
// 	dp := make([]int, 2*k)
// 	// 第一天只有有持有股票，收益一定为 -prices[0]，不管当天交易买入卖出多少次
// 	for i := 0; i < 2*k; i++ {
// 		if i&1 == 0 {
// 			dp[i] = -prices[0]
// 		}
// 	}
// 	for i := 1; i < n; i++ {
// 		for j := 2*k - 1; j > 0; j-- {
// 			temp := prices[i]
// 			// 如果是买入，第 k 次持有，就将价格改为负数，当 j 为偶数时，都是持有状态
// 			if j&1 == 0 {
// 				temp = -temp
// 			}
// 			dp[j] = max(dp[j], dp[j-1]+temp)
// 		}
// 		dp[0] = max(dp[0], -prices[i])
// 	}
// 	return dp[2*k-1]
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
