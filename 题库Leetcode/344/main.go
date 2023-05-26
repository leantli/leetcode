package main

// https://leetcode.cn/problems/reverse-string/
// 344. 反转字符串

// 两个指针分别指向首尾，swap 内容后，各自向中间步进
func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
