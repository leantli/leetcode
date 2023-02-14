package main

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
// 3. 无重复字符的最长子串

// 最长子串 不含重复字符
// 不定长滑动窗口
// 窗口内无重复字符
// 每次窗口 r 右移扩大时，看看新进的字符上次出现的位置，其是否在窗口内，是则缩小l
// l 直接缩小到 新进字符上次出现的位置右侧(+1)，仍保证窗口内无重复字符
// 并且注意每次都要更新 r 指向字符的最新位置
// 以及每次确认好窗口后看看窗口长度，记录最长长度
func lengthOfLongestSubstring(s string) int {
	bs := []byte(s)
	var l, r, maxLen int
	last := make(map[byte]int)
	for r < len(s) {
		// 看看这个字符之前有没有出现过
		// 如果出现过，是不是在当前窗口内(大于等于 l)
		// 如果是，l 右移至该字符上次出现的下标+1，此时窗口内仍然无重复字符
		// 如果不在窗口内，就继续 r 右移扩大窗口
		if index, ok := last[bs[r]]; ok {
			if index >= l {
				l = index + 1
			}
		}
		// 每次都要记录当前字符最后一次出现的位置
		last[bs[r]] = r
		r++
		maxLen = max(maxLen, r-l)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
