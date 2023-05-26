package main

import "sort"

// https://leetcode.cn/problems/assign-cookies/
// 455. 分发饼干

// 我们尽量保证饼干的大小刚刚好满足孩子的胃口
// 先对二者排序
func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	var gi, cnt int
	for si := 0; si < len(s); si++ {
		if gi == len(g) {
			break
		}
		if s[si] >= g[gi] {
			cnt++
			gi++
		}
	}
	return cnt
}
