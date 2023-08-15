package main

// https://leetcode.cn/problems/word-break-ii/description/
// 140. 单词拆分 II

// // dfs 这道题用例可能存在问题，不然模板 dfs 不应该成为困难题，看题解是需要结合 139 的思路去 dp 解决，非常麻烦
// func wordBreak(s string, wordDict []string) []string {
//     // dfs wordDict 中每个词，当能够拼成 s 时，join 空格加入到res中
//     res := make([]string, 0)
//     cur := make([]string, 0)
//     var dfs func()
//     dfs = func() {
//         temp := strings.Join(cur,"")
//         if !strings.HasPrefix(s, temp) || len(temp) > len(s) {
//             // 当前面的都不匹配，直接返回
//             return
//         }
//         if s == temp {
//             res = append(res, strings.Join(cur, " "))
//             return
//         }
//         for i := 0; i < len(wordDict); i++ {
//             cur = append(cur, wordDict[i])
//             dfs()
//             cur = cur[:len(cur)-1]
//         }
//     }
//     dfs()
//     return res
// }

// 困难版的 dp，思路确实很难想到
// 之前 139 中，dp[j] 存的是 s[0:j] 匹配是否为 true，现在可以存放匹配成功的字符串，匹配失败则不存放
// dp[j] = append(dp[j], ( for _, sentence := range dp[j-len(word)] ) + " " + word), if s[j-len(word):j] == word
func wordBreak(s string, wordDict []string) []string {
	dp := make([][]string, len(s)+1)
	dp[0] = []string{""}
	for j := 1; j <= len(s); j++ {
		for _, word := range wordDict {
			if j >= len(word) && s[j-len(word):j] == word {
				// 当前的词匹配成功，再去看之前 dp[j-len(word)] 中包含的所有句子
				// 逐一添加到当前的 dp[j] 中
				for _, sentence := range dp[j-len(word)] {
					tmp := word
					if sentence != "" {
						tmp = sentence + " " + word
					}
					dp[j] = append(dp[j], tmp)
				}
			}
		}
	}
	return dp[len(s)]
}
