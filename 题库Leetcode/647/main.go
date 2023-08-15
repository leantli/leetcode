package main

// https://leetcode.cn/problems/palindromic-substrings/description/
// 647. 回文子串

// 显然这是一道单串 dp，但是这里我们不能使用 dp[i] 表示以 s[i] 为结尾的字符串的回文子串数量
// 因为我们显然需要借助回文子串左右两侧的字符是否相等来判断新的位置是否是回文子串
// 所以这里显然需要使用带一维状态的 dp
// 定义 dp[i][j] 表示 s[i:j+1] 是不是一个回文子串，是为 true
// 则 显然 dp[i][j] = s[i] == s[j] && ( dp[i+1][j-1] || j-i < 2 )
// 并且我们可以关注到 dp[i][j] 由 dp[i+1][j-1] 决定，是自高往低的，因此我们在遍历时 i 取值从顶处开始
// 而 j 显然是子串的右区间，因此 j 必定大于等于 i
func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	var res int
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i < 2 || dp[i+1][j-1]) {
				res++
				dp[i][j] = true
			}
		}
	}
	return res
}

// // 使用中心拓展法进行统计计算
// func countSubstrings(s string) int {
// 	cnt := getCnt(0, 0, s)
// 	for i := 1; i < len(s); i++ {
// 		cnt += getCnt(i-1, i, s) + getCnt(i, i, s)
// 	}
// 	return cnt
// }
// func getCnt(l, r int, s string) int {
// 	var cnt int
// 	for l >= 0 && r < len(s) && s[l] == s[r] {
// 		l--
// 		r++
// 		cnt++
// 	}
// 	return cnt
// }
