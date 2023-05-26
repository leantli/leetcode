package main

// https://leetcode.cn/problems/decode-ways/
// 91. 解码方法

// 把刚才的分类重新进行简化
// 从上面我们知道，假定 dp[i] 为以 s[i] 结尾的字符串的解码情况数量
// 当 s[i] 单独成字符时，dp[i] = dp[i-1]，s[i] != '0' 时都可以单独成字符
// 当 s[i] 和 s[i-1] 能组成 10-26 的字符时，dp[i] = dp[i-2]
// 这里我们可以通过判断累加，简单实现区分，而不需要麻烦的 if elseif else
// 并且，我们注意到，i 只和 i-1、i-2 相关，因此我们可以用两个局部变量替代 dp 数组的声明
func numDecodings(s string) int {
	ppre, pre := 0, 1
	for i := range s {
		var cur int
		if s[i] != '0' {
			cur += pre
		}
		if i > 0 && s[i-1] != '0' && (s[i-1]-'0')*10+s[i]-'0' <= 26 {
			cur += ppre
		}
		ppre, pre = pre, cur
	}
	return pre
}

// // 1 - A
// // 11 - A A || K
// // 111 - (A A A || K A [11 + 1]) || (A K [1 + 11])
// // 从上面我们知道，假定 dp[i] 为以 s[i] 结尾的字符串的解码情况数量
// // 当 s[i] 单独成字符时，dp[i] = dp[i-1]，s[i] != '0' 时都可以单独成字符
// // 当 s[i] 和 s[i-1] 能组成 10-26 的字符时，dp[i] = dp[i-2]
// // 当 s[i] 同时满足两个时，dp[i] = dp[i-1] + dp[i-2]
// // 并且，我们注意到，i 只和 i-1、i-2 相关，因此我们可以用两个局部变量替代 dp 数组的声明
// func numDecodings(s string) int {
// 	ppre, pre := 0, 1
// 	for i := range s {
// 		var cur int
// 		if s[i] != '0' {
// 			cur += pre
// 		}
// 		if i > 0 && s[i] != '0' && s[i-1] != '0' && (s[i-1]-'0')*10+s[i]-'0' <= 26 {
// 			ppre, pre = pre, ppre+pre
// 		} else if i > 0 && s[i-1] != '0' && (s[i-1]-'0')*10+s[i]-'0' <= 26 {
// 			ppre, pre = pre, ppre
// 		} else if s[i] != '0' {
// 			ppre = pre
// 		} else {
// 			ppre, pre = pre, 0
// 		}
// 	}
// 	return pre
// }
