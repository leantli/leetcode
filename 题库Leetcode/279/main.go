package main

import (
	"math"
)

// https://leetcode.cn/problems/perfect-squares/description/
// 279. 完全平方数

// 有点类似零钱兑换，求和为 n 的最少数量
// 但是这里是物品是完全平方数，还需要用户自行先计算出来
// 接着是完全背包问题，求最少数量的方程为 dp[i][j] = min(dp[i-1][j], dp[i-1][j-nums[i]]+1)
// 滚动数组优化降维后为 dp[j] = min(dp[j], dp[j-nums[i]]+1)
func numSquares(n int) int {
	// 1. 求出完全平方数，作为可选取的物品
	// 由于 n >= 1 && n <= 10000，所以最后一个完全平方数为 100 即可
	// 并且其实考虑到，n 可能还不到 100，其实可选取的数，仅仅只是 1~sqrt(n)，而不需要1~100
	num := int(math.Sqrt(float64(n)))
	bags := make([]int, 0, num)
	for i := 1; i <= num; i++ {
		bags = append(bags, i*i)
	}
	// 2. 开始完全背包
	// 和为 n，则认定背包容量上限为 n
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = 10001
	}
	dp[0] = 0
	for _, bag := range bags {
		for j := 0; j <= n; j++ {
			if j >= bag {
				dp[j] = min(dp[j], dp[j-bag]+1)
			}
		}
	}
	return dp[n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
