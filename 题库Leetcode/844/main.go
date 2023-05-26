package main

// https://leetcode.cn/problems/backspace-string-compare/
// 844. 比较含退格的字符串

// 从后往前即可
// 遇到 # 就存储一次越过机会
// 当有越过机会时优先使用
// 无越过机会时再比较此时两个指针指向的字母
func backspaceCompare(s string, t string) bool {
	si, ti := len(s)-1, len(t)-1
	var sj, tj int
	for si >= 0 || ti >= 0 {
		for si >= 0 && (s[si] == '#' || sj > 0) {
			if s[si] != '#' {
				sj--
			} else if s[si] == '#' {
				sj++
			}
			si--
		}
		for ti >= 0 && (t[ti] == '#' || tj > 0) {
			if t[ti] != '#' {
				tj--
			} else if t[ti] == '#' {
				tj++
			}
			ti--
		}
		// 此时 si 和 ti 要么都到了某一个字母
		// 要么两个都刚好越界，如果只有一个越界那说明一定不相等
		if si >= 0 && ti >= 0 {
			if s[si] != t[ti] {
				return false
			}
		} else if si >= 0 || ti >= 0 {
			return false
		}
		si--
		ti--
	}
	return true
}
