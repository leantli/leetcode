package main

// https://leetcode.cn/problems/longest-valid-parentheses/
// 32. 最长有效括号

// 二刷
// 求最长的连续有效括号的长度
// 设 dp[i] 为以 i 下标为结尾的右括号，最长有效括号长度为多少
// 当 s[i] 为左括号时，dp[i] = 0
// 当 s[i] 为右括号时，若 s[i-1] = '('，则 dp[i] = dp[i-2] + 2
// 当 s[i] 为右括号时，若 s[i-1] = ')'，则查看 s[i-dp[i-1]-1] 是否为 '('，是则 dp[i] = dp[i-1]+2+dp[i-dp[i-1]-2]，否则 dp[i] = 0
// 这里为什么要加 dp[i-dp[i-1]-2]？比如说 ()(())，当我们遍历到最后一个右括号时，其和下标为 2 的 ( 匹配，但是最长有效括号并不只是 4，而是 6
// 因为最左侧还有正常的有效的 () 需要计入
func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	var res int
	for i := range s {
		if s[i] == '(' {
			continue
		}
		if i-1 >= 0 && s[i-1] == '(' {
			dp[i] = 2
			if i-2 >= 0 {
				dp[i] += dp[i-2]
			}
		} else if i-1 >= 0 && s[i-1] == ')' {
			if i-dp[i-1]-1 < 0 || s[i-dp[i-1]-1] == ')' {
				continue
			}
			dp[i] = dp[i-1] + 2
			if i-dp[i-1]-2 >= 0 {
				dp[i] += dp[i-dp[i-1]-2]
			}
		}
		res = max(res, dp[i])
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// // 比较巧妙的做法，不过我想不到就是了(
// // 利用两个计数器 l、r，遍历时计数左右括号的数量
// // 当 l, r 计数相同时，计算他们的长度，此时左右括号相等，能够正常匹配
// // 当 r 大于 l 时，l 和 r 都归零，因为右括号多了，没有多余的左括号匹配
// // 但是这样的话 (() 这样的括号就无法计算出最大长度
// // 因此我们可以反着再来一次，从尾部向前遍历，当 l 大于 r 时两个计数器都归零
// func longestValidParentheses(s string) int {
// 	var l, r, res int
// 	for i := range s {
// 		if s[i] == '(' {
// 			l++
// 		} else {
// 			r++
// 		}
// 		if l == r {
// 			res = max(l+r, res)
// 		} else if r > l {
// 			l, r = 0, 0
// 		}
// 	}
// 	l, r = 0, 0
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if s[i] == ')' {
// 			r++
// 		} else {
// 			l++
// 		}
// 		if l == r {
// 			res = max(l+r, res)
// 		} else if r < l {
// 			l, r = 0, 0
// 		}
// 	}
// 	return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// // 找出最长格式正确且连续的括号子串
// // 格式正确的话最后结尾的一定是右括号
// // 考虑最终目标是求长度为 n 的字符串 s 的最长有效括号
// // 我们可以拆解为求长度为 1~n 的字符串的最长有效括号
// // 或者求以 s[i] 为结尾时的最长有效括号的长度, 定义为 dp[i]
// // 那我们知道当 s[i] == ')' 且 s[i-1] == '(' 时，dp[i] = dp[i-2]+2
// // 当 s[i] == ')' 且 s[i-1] == ')' 时，那么我们要考虑 s[i-1] 的这个右括号，对应的左括号的左侧，是否还有一个左括号
// // 对应着 s[i] 的右括号，即 s[i-dp[i-1]] 就是 s[i-1] 右括号对应的左括号的下标，此时 s[i-dp[i-1]-1] 就是 s[i] 对应左括号应该在的位置
// // 如果 s[i-dp[i-1]-1] = '('，则说明又成了一对新括号，那么是不是就直接是 dp[i] = dp[i-1] + 2 了呢？
// // 这里其实还忽略了一种情况，如果这新成的一整块有效括号的左侧，也是一个有效括号呢？此时就应该延长
// // 所以应该是 dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
// // 而当 s[i] == '('，此时dp[i]=0，因为左括号结尾一定不会是有效括号
// func longestValidParentheses(s string) int {
// 	dp := make([]int, len(s))
// 	var res int
// 	for i := 1; i < len(s); i++ {
// 		if s[i] == ')' {
// 			if s[i-1] == '(' {
// 				if i >= 2 {
// 					dp[i] = dp[i-2] + 2
// 				} else {
// 					dp[i] = 2
// 				}
// 			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
// 				if i-dp[i-1]-2 >= 0 {
// 					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
// 				} else {
// 					dp[i] = dp[i-1] + 2
// 				}
// 			}
// 		}
// 		if res < dp[i] {
// 			res = dp[i]
// 		}
// 	}
// 	return res
// }

// // 看起来很像 20 题的进阶题，首先我们会想到还是用栈？只是少了括号的类型，更注重匹配后的最长长度，
// // 那么显然除了匹配之外，需要存储左括号的下标位置，遇到右括号匹配上时，才能用当前长度减去左括号的下标位置+1取得真正长度
// // 但会发现，遇到 )()()这样的就无法计算出最大长度，只能计算出 2 而不是 4
// // 因此当栈为空时，我们存储最后一个没匹配上的括号的下标, 并且考虑到整个字符串都是有效括号的情况，我们需要在栈中先存入一个-1，
// // -1 用于处理未出现未匹配括号的情况，比如 ()()
// func longestValidParentheses(s string) int {
// 	// 栈存储左括号的下标
// 	stack := []int{-1}
// 	var res int
// 	for i := range s {
// 		if s[i] == '(' {
// 			stack = append(stack, i)
// 			continue
// 		}
// 		// 遇到右括号，就将栈中左括号/上一个最后未匹配到的值 弹出
// 		stack = stack[:len(stack)-1]
// 		// 当栈中啥都没有，则表示未匹配成功，将这个括号的下标作为最后一个未匹配到的压入栈中
// 		if len(stack) == 0 {
// 			stack = append(stack, i)
// 			continue
// 		}
// 		dis := i - stack[len(stack)-1]
// 		if res < dis {
// 			res = dis
// 		}
// 	}
// 	return res
// }
