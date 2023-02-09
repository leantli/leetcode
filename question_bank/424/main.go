package main

// https://leetcode.cn/problems/longest-repeating-character-replacement/
// 424. 替换后的最长重复字符

// 不定长滑动窗口，以窗口首位作为 窗口内趋同的字母，k 为窗口内容许的非趋同字母数量
// 当窗口内所有字母都趋同时，记录此时长度，并且右移
// 不满足条件时，l 右移
func characterReplacement(s string, k int) int {
	n := len(s)
	if k == n {
		return n
	}
	var l, r int
	bs := []byte(s)
	cMap := make([]int, 128)
	var maxLen, maxCount int // 最大长度 和 窗口内出现频率最高字母的次数
	for r < n {
		cMap[bs[r]]++
		maxCount = max(maxCount, cMap[bs[r]])
		r++
		// r-l 即窗口内的字符数量，
		// 如果 出现频率最高字母的次数 + k 小于该数量，则说明此时窗口内字母无法趋同
		// 此时 l 右移
		for r-l > maxCount+k {
			cMap[bs[l]]--
			l++
		}
		maxLen = max(r-l, maxLen)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
