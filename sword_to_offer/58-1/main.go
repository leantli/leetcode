package main

import "strings"

// https://leetcode.cn/problems/fan-zhuan-dan-ci-shun-xu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 58-1. 翻转单词顺序

func reverseWords(s string) string {
	bs := []byte(s)
	n := len(bs)
	temp := make([][]byte, 0, n)
	l := 0
	for r := 0; r < n; r++ {
		for r < n && bs[r] == ' ' {
			r++
		}
		l = r
		for r < n && bs[r] != ' ' {
			r++
		}
		if l == r && l == n {
			break
		}
		temp = append(temp, bs[l:r])
	}
	var ts []string
	for i := len(temp) - 1; i >= 0; i-- {
		ts = append(ts, string(temp[i]))
	}
	return strings.Join(ts, " ")
}
