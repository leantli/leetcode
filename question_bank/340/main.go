package main

// https://leetcode.cn/problems/longest-substring-with-at-most-k-distinct-characters/
// 340. 至多包含 K 个不同字符的最长子串

// 159 的小进阶
// 窗口中维护一个字符数组，每次有从0到1的，窗口内字符数+1，有变成0的，窗口内字符数就-1
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	var charCnt [128]int
	var winCnt int
	var l, r, maxLen int
	for r < len(s) {
		c := s[r]
		if charCnt[c] == 0 {
			winCnt++
		}
		charCnt[c]++
		for winCnt > k {
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
