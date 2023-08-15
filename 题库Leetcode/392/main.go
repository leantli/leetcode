package main

// https://leetcode.cn/problems/is-subsequence/description/
// 392. 判断子序列

// 但这道题同时也非常类似双串dp，定义 dp[i][j] 为 s[0:i] 与 t[0:j] 公共子序列长度
// 则 dp[i][j] = dp[i-1][j-1]+1, if s[i-1] == t[i-1]
// dp[i][j] = dp[i][j-1], if s[i-1] != t[j-1]，因为这里是判断 s 是否为 t 的子序列，
// 不需要考虑 s 长 t 短的情况，必定是 t 更长，因此 dp[i][j] 不选取 dp[i-1][j]，只选取 dp[i][j-1]
// 当然，加上 max 考虑也可以，dp[i][j] = max(dp[i][j-1], dp[i-1][j]), if s[i-1] != t[j-1]
// 最后判断 dp[n][m] == n 即可
func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[n][m] == n
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// // 判断一个字符串是否是另一个字符串的子序列
// func isSubsequence(s string, t string) bool {
// 	if len(s) > len(t) {
// 		return false
// 	}
// 	// s 一定比 t 长度短，s 为 t 子序列的话，首先各个字母需要依序在 t 中出现
// 	// 依靠依序在 t 中出现，我们可以基于两个指针不断进行比较
// 	// 分别指向 s 和 t 的两个指针
// 	var i, j int
// 	for i < len(s) && j < len(t) {
// 		for i < len(s) && j < len(t) && s[i] == t[j] {
// 			i++
// 			j++
// 		}
// 		j++
// 	}
// 	if i == len(s) {
// 		return true
// 	}
// 	return false
// }
