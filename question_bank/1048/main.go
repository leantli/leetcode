package main

import "sort"

// https://leetcode.cn/problems/longest-string-chain/
// 1048. 最长字符串链

// 这道题和 LIS 基本没有区别
// 唯一要注意的是，我们需要执行实现 a 是否为 b 前身的函数
// 其次要先基于字符串长度排序一次，便于前身的判断
func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) <= len(words[j]) })
	// dp[i] 为 以words[i]结尾时词链的最长可能长度
	dp := make([]int, len(words))
	// 初始化，单个词自成最长词链
	for i := range words {
		dp[i] = 1
	}
	longest := 1
	for i := 1; i < len(words); i++ {
		for j := 0; j < i; j++ {
			if isAPreB(words[j], words[i]) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		longest = max(longest, dp[i])
	}
	return longest
}

// 判断 a 是不是 b 的前身
func isAPreB(a, b string) bool {
	if len(a)+1 != len(b) {
		return false
	}
	var chance bool
	for ai, bi := 0, 0; ai < len(a); {
		if a[ai] != b[bi] {
			if chance {
				return false
			}
			chance = true
			bi++
			continue
		}
		ai, bi = ai+1, bi+1
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(isAPreB("a", "ac"))
// 	fmt.Println(isAPreB("a", "acd"))
// 	fmt.Println(isAPreB("a", "ba"))
// 	fmt.Println(isAPreB("acd", "acdq"))
// 	fmt.Println(isAPreB("cba", "bcad"))
// }
