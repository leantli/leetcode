package main

// https://leetcode.cn/problems/find-k-length-substrings-with-no-repeated-characters/
// 1100. 长度为 K 的无重复字符子串

// 返回长度为 k 且不含重复字符的子串数量
// 显然定长滑动窗口
// 长度为 k(定长)，窗口性质-不含重复字符
// 每次右扩就判断当前字符是否已在窗口中出现过，是则 l 移动到出现过的位置+1(右侧)，无则正常继续操作
// 此时再判断当前长度是否超过 k，是则抛弃最左边的，再判断是否长度等于 k，是则 res++
func numKLenSubstrNoRepeats(s string, k int) int {
	n := len(s)
	if k > n {
		return 0
	}
	var res int
	var r, l int
	last := make(map[byte]int) // 存储字符-上次出现的位置
	for r < n {
		if index, ok := last[s[r]]; ok {
			if index >= l {
				l = index + 1
			}
		}
		if r-l+1 > k {
			l++
		}
		if r-l+1 == k {
			res++
		}
		last[s[r]] = r
		r++
	}
	return res
}
