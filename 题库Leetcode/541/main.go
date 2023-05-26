package main

// https://leetcode.cn/problems/reverse-string-ii/
// 541. 反转字符串 II

// 翻转 k 个字符 [l,l+k-1]
// 然后到下一个翻转区域的起始点 l+2k
// 当要翻转区域的 tr 超过了字符串长度，则取到末尾即可
func reverseStr(s string, k int) string {
	bs := []byte(s)
	for l := 0; l < len(s); l += 2 * k {
		tl, tr := l, l+k-1
		if tr >= len(s) {
			tr = len(s) - 1
		}
		for tl < tr {
			bs[tl], bs[tr] = bs[tr], bs[tl]
			tl++
			tr--
		}
	}
	return string(bs)
}
