package main

// https://leetcode.cn/problems/longest-palindromic-substring/submissions/
// 5. 最长回文子串

// 一道单串二维dp
// 定义 dp[i][j] 为 s[i:j+1] 是否为回文子串，是则为 true
// 显然 dp[i][j] =  s[i] == s[j] && (dp[i+1][j-1] || j-i < 2 )
// 当 dp[i][j] == true 时，再统计当前的左右边界，选取回文子串最长的边界进行记录，方便最终返回
// 初始化 dp[i][i] 显然都为 true，自身成回文子串
// 由于 dp[i][j] 由 dp[i+1][j-1] 决定，我们的 i 遍历，需要从大到小，j 则正常从小到大
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	var resL, resR int
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] && (dp[i+1][j-1] || j-i < 2) {
				dp[i][j] = true
				if j-i > resR-resL {
					resR = j
					resL = i
				}
			}
		}
	}
	// dp[i][j] 表示 s[i:j+1]，因此记录的 resL 和 resR，resR 需要加以处理，保证返回闭区间
	return s[resL : resR+1]
}

// // 中心拓展法找最长回文子串
// func longestPalindrome(s string) string {
// 	resL, resR := getLenght(0, 0, s)
// 	for i := 1; i < len(s); i++ {
// 		temp1L, temp1R := getLenght(i, i, s)
// 		temp2L, temp2R := getLenght(i-1, i, s)
// 		if resR-resL < temp1R-temp1L {
// 			resL = temp1L
// 			resR = temp1R
// 		}
// 		if resR-resL < temp2R-temp2L {
// 			resL = temp2L
// 			resR = temp2R
// 		}
// 	}
// 	return s[resL+1 : resR]
// }
// func getLenght(l, r int, s string) (int, int) {
// 	var length int
// 	for l >= 0 && r < len(s) && s[l] == s[r] {
// 		length++
// 		l--
// 		r++
// 	}
// 	return l, r
// }
