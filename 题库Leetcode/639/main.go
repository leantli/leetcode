package main

// https://leetcode.cn/problems/decode-ways-ii/
// 639. 解码方法 II

// // 91 的进阶题，多出了 * 字符的处理，* 可以对应 1-9 任意一个数处理
// // 定义 dp[i] 为以 s[i] 结尾的字符串的解码情况数量
// // 当 s[i] 单独成字符时, 只要 s[i] 不为 0，dp[i] = dp[i-1]，并且当 s[i] 为 '*' 时，dp[i] = 9*dp[i-1]
// // 当 s[i] 和 s[i-1] 能组成字符时，只要 s[i-1] * 10 + s[i] >= 10 并且 <= 26，则 dp[i] = dp[i-2]
// // 当 s[i] 为 '*' 并且和 s[i-1] 组成字符时，当 s[i-1] 为 '1' 时，dp[i] = 9*dp[i-2]
// // 当 s[i-1] 为 2 时，dp[i] = 6*dp[i-2]，其他情况都无需考虑
// // 当 s[i-1] 和 s[i] 都为 '*' 时，dp[i] = 15*dp[i-2]
// // 当 s[i-1] 为 '*'，0 <= s[i] <=6 时 dp[i] = 2*dp[i-2]
// // 当 s[i-1] 为 '*'，s[i] >= 7 时 dp[i] = dp[i-2]
// // 并且我们注意到 dp[i] 只和 i-1, i-2 相关，因此我们可以用两个局部变量替代dp数组的声明
// func numDecodings(s string) int {
// 	mod := int(1e9 + 7)
// 	ppre, pre := 0, 1
// 	for i := range s {
// 		var temp int
// 		if s[i] != '*' {
// 			if s[i] != '0' {
// 				temp = (temp + pre) % mod
// 			}
// 			if i > 0 && s[i-1] != '*' && s[i-1] != '0' && (s[i-1]-'0')*10+s[i]-'0' <= 26 {
// 				temp = (temp + ppre) % mod
// 			}
// 			if i > 0 && s[i-1] == '*' && s[i] <= '6' {
// 				temp = (temp + 2*ppre) % mod
// 			}
// 			if i > 0 && s[i-1] == '*' && s[i] >= '7' {
// 				temp = (temp + ppre) % mod
// 			}
// 		} else {
// 			temp = (9 * pre) % mod
// 			if i > 0 && s[i-1] == '1' {
// 				temp = (temp + 9*ppre) % mod
// 			}
// 			if i > 0 && s[i-1] == '2' {
// 				temp = (temp + 6*ppre) % mod
// 			}
// 			if i > 0 && s[i-1] == '*' {
// 				temp = (temp + 15*ppre) % mod
// 			}
// 		}
// 		ppre, pre = pre, temp
// 	}
// 	return pre
// }

// 佬的写法
func check1digit(ch byte) int {
	if ch == '*' {
		return 9
	}
	if ch == '0' {
		return 0
	}
	return 1
}

func check2digits(c0, c1 byte) int {
	if c0 == '*' && c1 == '*' {
		return 15
	}
	if c0 == '*' {
		if c1 <= '6' {
			return 2
		}
		return 1
	}
	if c1 == '*' {
		if c0 == '1' {
			return 9
		}
		if c0 == '2' {
			return 6
		}
		return 0
	}
	if c0 != '0' && (c0-'0')*10+(c1-'0') <= 26 {
		return 1
	}
	return 0
}

func numDecodings(s string) int {
	const mod int = 1e9 + 7
	a, b, c := 0, 1, 0
	for i := range s {
		c = b * check1digit(s[i]) % mod
		if i > 0 {
			c = (c + a*check2digits(s[i-1], s[i])) % mod
		}
		a, b = b, c
	}
	return c
}
