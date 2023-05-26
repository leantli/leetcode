package main

// https://leetcode.cn/problems/longest-substring-with-at-most-two-distinct-characters/
// 159. 至多包含两个不同字符的最长子串

// 更简洁一点的思路
// 窗口中维护一个字符数组，每次有从0到1的，窗口内字符数+1，有变成0的，窗口内字符数就-1
func lengthOfLongestSubstringTwoDistinct(s string) int {
	var charCnt [128]int
	var winCnt int
	var l, r, maxLen int
	for r < len(s) {
		c := s[r]
		if charCnt[c] == 0 {
			winCnt++
		}
		charCnt[c]++
		for winCnt > 2 {
			lc := s[l]
			if charCnt[lc] == 1 {
				winCnt--
			}
			charCnt[lc]--
			l++
		}
		r++
		if r-l > maxLen {
			maxLen = r - l
		}
	}
	return maxLen

}

// // 最长子串 不定长滑窗？
// // 窗口性质-窗口内最多包含两个不同字符
// // 用两个 byte 记录当前窗口已经包含的字符，和两个 int 记录这两个字符在窗口中最后出现的位置
// // 每当遇到不属于这两个字符的，就更新其中一个字符
// // 那么更新哪个字符？更新在窗口中更靠左的那个字符，比如说"aabbcc"，我们肯定是把 a 换成 c
// // 并且窗口的左边界 l 也要移到a 最后一次出现位置的右侧
// func lengthOfLongestSubstringTwoDistinct(s string) int {
// 	var l, r, maxLen int
// 	// 初始化窗口
// 	// 窗口需要维护 l,r 两个边界
// 	// 以及窗口内是哪两个字符，以及这两个字符上次出现的最后位置
// 	var a, b byte
// 	var aIndex, bIndex int
// 	a = s[0]
// 	for r < len(s) {
// 		if s[r] != a {
// 			b = s[r]
// 			bIndex = r
// 			r++
// 			break
// 		}
// 		aIndex = r
// 		r++
// 	}
// 	maxLen = r - l
// 	for r < len(s) {
// 		// 先进来一个数，然后看看是否破坏窗口性质，是则维护窗口性质
// 		if s[r] != a && s[r] != b {
// 			// 此时破坏窗口性质，移动窗口左边界到两个字符中更靠左的字符最后一次出现位置的右侧
// 			// 并更新窗口中的两个字符和位置
// 			fmt.Printf("aI:%d, bI:%d\n", aIndex, bIndex)
// 			if bIndex < aIndex {
// 				l = bIndex + 1
// 				b = s[r]
// 				bIndex = r
// 			} else {
// 				l = aIndex + 1
// 				a = s[r]
// 				aIndex = r
// 			}
// 		} else if s[r] == a {
// 			aIndex = r
// 		} else if s[r] == b {
// 			bIndex = r
// 		}
// 		r++
// 		fmt.Printf("l:%d, r:%d\n", l, r)
// 		if r-l > maxLen {
// 			maxLen = r - l
// 		}
// 	}
// 	return maxLen
// }
