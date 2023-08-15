package main

// https://leetcode.cn/problems/longest-palindromic-subsequence/description/
// 516. 最长回文子序列

// 注意，这里是最长回文子序列而不是最长回文子串，因此我们不方便使用暴力或中心拓展法等方式
// 考虑 dp 解决，而这显然是单串 dp 但带有一维状态
// 定义 dp[i][j] 为 s[i:j+1] 子串中最长的回文子序列长度
// if s[i] == s[j], dp[i][j] = dp[i+1][j-1] + 2, else dp[i][j] = max(dp[i+1][j], dp[i][j-1])
// s[i] == s[j] 时，i 和 j 相当于都能取，s[i] != s[j] 时，显然 i 和 j 只能取其一，选择回文子串长度最长的取
// 同时我们关注到，dp[i][j] 只由 dp[i+1][j-1]、dp[i+1][j]、dp[i][j-1] 决定，都是 i、j 左上侧的
// 因此 i 我们需要从顶往底遍历，而 j 左右右区间，显然 j >= i
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1 // 初始化，默认情况下，dp[i][j] 显然是单独成回文子串的，长度为 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
