package main

// https://leetcode.cn/problems/number-of-substrings-containing-all-three-characters/
// 1358. 包含所有三种字符的子字符串数目

// 不定长滑窗
// 窗口性质--窗口中a,b,c都存在
// 不满足时右移右边界
// 满足时先计算基于当前窗口，有多少数目，再右移左边界
// 相同字符串可算多次
func numberOfSubstrings(s string) int {
	n := len(s)
	var res int
	var l, r, curCount int // 窗口当前有多少种字符
	var cnt [3]int
	for r < n {
		c := s[r] - 'a'
		cnt[c]++
		if cnt[c] == 1 {
			curCount++
		}
		// 满足条件时进行计算，以及移动窗口的左边界
		for curCount == 3 {
			res += n - r
			lc := s[l] - 'a'
			cnt[lc]--
			if cnt[lc] == 0 {
				curCount--
			}
			l++
		}
		r++
	}
	return res
}
