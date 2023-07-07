package main

import "strings"

// https://leetcode.cn/problems/reverse-words-in-a-string/
// 151. 反转字符串中的单词

func reverseWords(s string) string {
	strs := make([]string, 0)
	var l, r int
	n := len(s)
	// 遍历，左右指针，第一次右指针遇到非空格时，将 l=r，r 再继续遍历，遇到空格时停下，此时 s[l:r] 就是一个单词
	for r < len(s) {
		for r < n && s[r] == ' ' {
			r++
		}
		l = r
		for r < n && s[r] != ' ' {
			r++
		}
		if l == r && l == n {
			break
		}
		strs = append(append([]string{}, s[l:r]), strs...)
	}
	return strings.Join(strs, " ")
}
