package main

import "math"

// https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 63. 股票的最大利润

// 属实忘记 dp 的状态转移方程了，偷偷看了一眼，看看能不能搞出来
// dp[i] = max{ dp[i-1], price[i]-min }
func maxProfit(prices []int) int {
	minPrice := math.MaxInt // 前i天最低的价格
	profit := 0             // 当天卖出可获得的最大利润
	for _, price := range prices {
		if minPrice > price {
			minPrice = price
		}
		profit = max(profit, price-minPrice)
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// // 继续先不用 dp，借用个 map，只保存 截止到遍历位置前 最低点股票的价格
// // 这样每到新的位置，就用当前价格去减之前最低的价格
// // 再判断是否要更新当前的最低价格
// func maxProfit(prices []int) int {
// 	res := 0
// 	if len(prices) < 2 {
// 		return res
// 	}
// 	m := make(map[int]int)
// 	m[0] = prices[0]
// 	for i := 1; i < len(prices); i++ {
// 		d := prices[i] - m[i-1]
// 		if d > res {
// 			res = d
// 		}
// 		if prices[i] < m[i-1] {
// 			m[i] = prices[i]
// 		} else {
// 			m[i] = m[i-1]
// 		}
// 	}
// 	return res
// }

// 凭借着印象瞎做出来的，动态规划的状态转移方程忘记咋推了，把这个当作最大子序和来做了。。
// 初始化以第一天的价格买入
// 后面遍历，遍历的同时观察利润是否为正，是的话买入的时间不变，只判断是否是最大利润
// 若利润为负，就直接舍去，取当天为新的买入时间
// 但是整个推导过程属实忘了。。。
// func maxProfit(prices []int) int {
// 	res := 0
// 	if len(prices) < 2 {
// 		return res
// 	}
// 	buy := prices[0]
// 	for _, price := range prices {
// 		d := price - buy
// 		if d >= 0 {
// 			if d > res {
// 				res = d
// 			}
// 		} else {
// 			buy = price
// 		}
// 	}
// 	return res
// }

// 暴力，不写了
