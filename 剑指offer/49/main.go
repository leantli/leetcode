package main

// https://leetcode.cn/problems/chou-shu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 49. 丑数

// dp(n) = dp[n/2] || dp[n/3] || dp[n/5]
// 或者说 dp[n] = min{dp[p2]*2, dp[p3]*3, dp[p5]*5}
// 因为只包含质因子 2、3 和 5 ，因此我们可以从 1 开始，取三个指针，分别乘三个质因子，得到的数再反复乘三个质因子
// 这样就能够保证得到的都是丑数
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
		// 注意不能使用 switch，否则当 x2,x3,x5 结果相同时，只会移动单个指针，导致丑数存入重复
		// switch dp[i] {
		// case x2:
		// 	p2++
		// case x3:
		// 	p3++
		// case x5:
		// 	p5++
		// }
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
