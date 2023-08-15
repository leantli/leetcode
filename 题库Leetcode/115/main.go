package main

// https://leetcode.cn/problems/distinct-subsequences/
// 115. 不同的子序列

// 子序列问题显然还是很难使用回溯或者dfs等方法进行枚举，考虑 dp
// 算是典型的双串 dp，dp[i][j] 表示 s[i-1] 为结尾的序列中出现 t[:j-1] 子序列的个数
// 此时 if s[i-1] == t[j-1], dp[i][j] = dp[i-1][j-1] + dp[i-1][j]，为什么不是 dp[i][j] = dp[i-1][j-1] 呢？
// 注意，我们的 dp[i][j] 的定义是，以 s[i-1] 为结尾的序列中出现 t[:j-1] 子序列的个数
// 也就是说，当 s[i-1] == t[j-1] 时，dp[i][j] 肯定包含了 dp[i-1][j-1] 的个数，但同时也包含了 s[i-1] != t[j-1] 的匹配数量
// 所以 dp[i][j] = dp[i-1][j-1] + dp[i-1][j], if s[i-1] == t[i-1]，其实这里就类似于 01 背包求方法数的考虑
// 现在有7个物品“babgbag", 依次放进size为3的背包"bag"中，但放进背包的条件是只有字符相等才能放，可以选择放与不放
// 放的话 dp[i][j] = dp[i-1][j-1]，不放的话,d[i][j] = dp[i-1][j]，这里放与不放的条件就是 s[i-1] 是否等于 t[j-1]
// 而 dp[i][j] 的方法数，显然是放与不放的总和，所以 dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
// 再基于滚动数组优化降维，dp[j] += dp[j-1]，从后往前遍历
func numDistinct(s string, t string) int {
	n, m := len(s), len(t)
	dp := make([]int, m+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := m; j > 0; j-- {
			if s[i-1] == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[m]
}

// // 二维版本，未经过滚动数组优化降维
// func numDistinct(s string, t string) int {
// 	n, m := len(s), len(t)
// 	dp := make([][]int, n+1)
// 	for i := range dp {
// 		dp[i] = make([]int, m+1)
// 		dp[i][0] = 1 // 初始化
// 	}
// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= m; j++ {
// 			if s[i-1] == t[j-1] {
// 				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
// 			} else {
// 				dp[i][j] = dp[i-1][j]
// 			}
// 		}
// 	}
// 	return dp[n][m]
// }
