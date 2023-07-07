package main

import (
	"math"
)

// https://leetcode.cn/problems/nge-tou-zi-de-dian-shu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 60. n个骰子的点数

// 二刷
// dp[i] 表示点数为 i 的情况的次数，但是显然这样是不合理的，因为我们不确定有多少个骰子
// 无法单纯基于 dp[i] 去进行状态转移
// dp[i][j] 表示 i 个骰子时，点数为 j 的出现次数
// 此时容易考虑 dp[i][j] = dp[i-1][j - (1~6)] 总和
// dp[1][1~6] = 1/6; dp[2][2] =
func dicesProbability(n int) []float64 {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n*6+1)
	}
	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}
	// i 表示当前是第 i 个骰子
	for i := 2; i <= n; i++ {
		// j 表示当前 i 个骰子的和的范围
		for j := i * 1; j <= i*6; j++ {
			// k 表示第 i 个骰子掷出来的点数
			for k := 1; k <= 6; k++ {
				if j-k > 0 && dp[i-1][j-k] > 0 {
					dp[i][j] += dp[i-1][j-k]
				}
			}
		}
	}
	sum := math.Pow(6, float64(n))
	res := make([]float64, 0)
	for i := n; i <= n*6; i++ {
		res = append(res, float64(dp[n][i])/sum)
	}
	return res
}

// // 这题的第一个思路是 暴力 dfs 填坑模拟, 每个骰子 6 种可能，不过显然，到 11 的时候大概率爆栈，并且有相当多计算是累赘的
// // 因此这里活用 dp
// // 状态定义：
// // dp[i][j] i 表示多少个骰子，j 表示各种点数和的出现次数
// // dp[i][j] = dp[i-1][j-1] + dp[i-1][j-2] + .... + dp[i-1][j-6]

// func dicesProbability(n int) []float64 {
// 	dp := make([][]int, n+1)
// 	for i := 1; i <= n; i++ {
// 		dp[i] = make([]int, i*6+1)
// 	}
// 	// 初始化 1 个骰子中，各个点数和出现的次数
// 	for j := 1; j <= 6; j++ {
// 		dp[1][j] = 1
// 	}
// 	// 状态转移
// 	for i := 1; i < n; i++ {
// 		for j := i; j <= 6*i; j++ {
// 			for k := 1; k <= 6; k++ {
// 				if dp[i][j] > 0 {
// 					dp[i+1][j+k] += dp[i][j]
// 				}
// 			}
// 		}
// 	}
// 	sum := math.Pow(6, float64(n))
// 	res := make([]float64, 0, 6*n+1)
// 	for i := n; i <= 6*n; i++ {
// 		res = append(res, float64(dp[n][i])/sum)
// 	}
// 	return res
// }

// 分析问题的状态时，不要分析整体，只分析最后一个阶段即可！因为动态规划问题都是划分为多个阶段的，各个阶段的状态表示都是一样，而我们的最终答案在就是在最后一个阶段。
// 对于这道题，最后一个阶段是什么呢？
// 通过题目我们知道一共投掷 n 枚骰子，那最后一个阶段很显然就是：当投掷完 n 枚骰子后，各个点数出现的次数
// 找出了最后一个阶段，那状态表示就简单了
// --->
// 首先用数组的第一维来表示阶段，也就是投掷完了几枚骰子。
// 然后用第二维来表示投掷完这些骰子后，可能出现的点数。
// 数组的值就表示，该阶段各个点数出现的次数。
// 所以状态表示就是这样的：dp[i][j] ，表示投掷完 i 枚骰子后，点数 j 的出现次数
// for (第 n+1 枚骰子的点数 i = 1; i <= 6; i ++) {
//     dp[n+1][j+i] += dp[n][j]
// }
