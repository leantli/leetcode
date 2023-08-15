package main

// https://leetcode.cn/problems/delete-operation-for-two-strings/description/
// 583. 两个字符串的删除操作

// 先求最长的公共子序列，再用各自的长度减去最长的公共子序列长度即可？
// 求两个子串的公共子序列，显然是一个双串dp，定义 dp[i][j] 为 word1 0~i-1 和 word2 0~j-1 闭区间中，最长的公共子序列长度
// if word1[i-1] == word2[j-1], dp[i][j] = dp[i-1][j-1] + 1
// if word1[i-1] != word2[j-1], dp[i][j] = max(dp[i-1][j], dp[i][j-1])，不等则有两个来源，选取较长的情况
func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return (n - dp[n][m]) + (m - dp[n][m])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
