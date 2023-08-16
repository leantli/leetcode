package main

// https://leetcode.cn/problems/word-break/
// 139. 单词拆分

// // dfs 估计会超时 // 部分测试用例超时
// func wordBreak(s string, wordDict []string) bool {
//     var dfs func(cur string)
//     var match bool
//     dfs = func(cur string) {
//         // 如果 s 的开头不是 cur 字符串（匹配失败）或已经匹配上了，则直接返回，不用匹配了
//         if len(cur) > len(s) || !strings.HasPrefix(s, cur) || match {
//             return
//         }
//         if s == cur {
//             match = true
//             return
//         }
//         for i := 0; i < len(wordDict); i++ {
//             dfs(cur+wordDict[i])
//         }
//     }
//     dfs("")
//     return match
// }

// 感觉在一定程度类似完全背包，背包容量就是 s，是否满足条件则是看拼接的是否匹配，并且这里存在排列的问题
// 显然应该是背包在外循环，wordDict 在内循环的完全背包
// 完全背包基础状态转移公式：dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
// 但是这里要求的是 s 是否能用 wordDict 中的单词拼接而出，所以显然是一个 bool 类型的状态转移
// 并且背包 dp[j] 能否拼接而出，则取决于 dp[i] 是否为 true && s[i:j] 存在于 wordDict 中
// 因此 dp[j] = dp[i] && m[s[i:j]]
// func wordBreak(s string, wordDict []string) bool {
//     m := make(map[string]bool)
//     for _, word := range wordDict {
//         m[word] = true
//     }
//     dp := make([]bool, len(s)+1)
//     dp[0] = true
//     for j := 1; j <= len(s); j++ {
//         for i := 0; i < j; i++ {
//             if dp[i] && m[s[i:j]] {
//                 dp[j] = true
//                 break
//             }
//         }
//     }
//     return dp[len(s)]
// }

// // 如果是求组成 s 的方式有多少种，则是 dp[j] += dp[j-len(wordDict[i])], dp[0] = 1
// func wordBreak(s string, wordDict []string) bool {
// 	dp := make([]int, len(s)+1)
// 	dp[0] = 1
// 	for j := 1; j <= len(s); j++ {
// 		for _, word := range wordDict {
// 			if j >= len(word) && s[j-len(word):j] == word {
// 				dp[j] += dp[j-len(word)]
// 			}
// 		}
// 	}
// 	return dp[len(s)] > 0
// }

// 第三次优化
// 背包就是 s[0:0~i]，背包就是 wordDict[i]
// 这里是判断，是否能够拼接成功，则 dp[i][j] 表示，从 0~i 中选取 wordDict，放入背包 j 中，能否拼接出对应的单词
// dp[i][j] = dp[i-1][j-len(wordDict[i])] && wordDict[i] == s[j-len(wordDict[i]): j]
// 滚动数组优化降维 dp[j] = dp[j-len(wordDict[i])] && wordDict[i] == s[j-len(wordDict[i]): j] && !dp[j]
// 即当 dp[j-len(word)] 为 true，并且 dp[j] 本身还是 false 时，并且 s[j-len(word):j] == word，此时 dp[j] = true
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	// 有排序，则背包在外循环，物品在内循环，保证某个背包的位置，都尝试放过所有物品，这样能遍历到所有排列
	for j := 1; j <= len(s); j++ {
		for _, word := range wordDict {
			if j >= len(word) && !dp[j] {
				dp[j] = dp[j-len(word)] && word == s[j-len(word):j]
			}
		}
	}
	return dp[len(s)]
}
