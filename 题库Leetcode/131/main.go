package main

// https://leetcode.cn/problems/palindrome-partitioning/
// 131. 分割回文串

// dp 先确定好回文串的情况，再 dfs 枚举
func partition(s string) [][]string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	// dp[i][j] = dp[i+1][j-1] && s[i]==s[j]
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] && (dp[i+1][j-1] || j-i < 2) {
				dp[i][j] = true
			}
		}
	}
	res := make([][]string, 0)
	cur := make([]string, 0)
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(s) {
			res = append(res, append([]string{}, cur...))
			return
		}
		for i := idx; i < n; i++ {
			str := s[idx : i+1]
			if dp[idx][i] {
				cur = append(cur, str)
				dfs(i + 1)
				cur = cur[:len(cur)-1]
			}
		}
	}
	dfs(0)
	return res
}

// // dfs 枚举分割位置并且查看是否为回文串
// func partition(s string) [][]string {
// 	n := len(s)
// 	res := make([][]string, 0)
// 	cur := make([]string, 0)
// 	var dfs func(idx int)
// 	dfs = func(idx int) {
// 		if idx == len(s) {
// 			res = append(res, append([]string{}, cur...))
// 			return
// 		}
// 		for i := idx; i < n; i++ {
// 			str := s[idx : i+1]
// 			if isPalin(str) {
// 				cur = append(cur, str)
// 				dfs(i + 1)
// 				cur = cur[:len(cur)-1]
// 			}
// 		}
// 	}
// 	dfs(0)
// 	return res
// }

// // 是回文返回 true
// func isPalin(str string) bool {
// 	l, r := 0, len(str)-1
// 	for l < r {
// 		if str[l] != str[r] {
// 			return false
// 		}
// 		l++
// 		r--
// 	}
// 	return true
// }
