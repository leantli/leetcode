package main

// https://leetcode.cn/problems/number-of-valid-words-for-each-puzzle/
// 1178. 猜字谜

// 没做出来，看了一下题解--二进制状态压缩，以后做这类型题再细究吧
func findNumOfValidWords(words []string, puzzles []string) []int {

}

// // 输出 puzzles[i] 作为谜面情况下，words 中有多少个能做 puzzles[i] 谜底
// // 首先就是先都遍历一遍看看 words 中是否有 puzzles[i] 的首字母？
// // 顺便看看 words 中有没有用 puzzles[i] 中之外的字符
// // 不过这样应该会超时吧，不然这种 单纯模拟 不会是困难题
// func findNumOfValidWords(words []string, puzzles []string) []int {
// 	res := make([]int, 0, len(puzzles))
// 	for _, puzzle := range puzzles {
// 		var pMap [26]int
// 		for j := 0; j < len(puzzle); j++ {
// 			pMap[puzzle[j]-'a']++
// 		}
// 		var count int
// 		for _, word := range words {
// 			var special, wrong bool
// 			// word 要成为 puzzle 的谜底，需要满足两个条件
// 			// 1. word 中包含 puzzle 字符串的第一个字符
// 			// 2. word 中不包含 puzzle 字符串中没有的字符
// 			for k := 0; k < len(word); k++ {
// 				// 包含 puzzle 字符串的第一个字符
// 				if !special && word[k] == puzzle[0] {
// 					special = true
// 				}
// 				if pMap[word[k]-'a'] == 0 {
// 					wrong = true
// 					break
// 				}
// 			}
// 			if !wrong && special {
// 				count++
// 			}
// 		}
// 		res = append(res, count)
// 	}
// 	return res
// }
