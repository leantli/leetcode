package main

// https://leetcode.cn/problems/chou-shu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 264. 丑数 II

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	i2, i3, i5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		num2, num3, num5 := dp[i2]*2, dp[i3]*3, dp[i5]*5
		minest := min(min(num2, num3), num5)
		if num2 == minest {
			i2++
		}
		if num3 == minest {
			i3++
		}
		if num5 == minest {
			i5++
		}
		dp[i] = minest
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
