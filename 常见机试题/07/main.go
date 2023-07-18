package main

import "fmt"

/**
题目描述:
近些年来，我国防沙治沙取得显著成果。某沙漠新种植N棵胡杨(编号1-N)，排成一排。一个月后，有M棵胡杨未能成活。
现可补种胡杨K棵，请问如何补种(只能补种，不能新种)，可以得到最多的连续胡杨树?
**/

func main() {
	fmt.Println(work(5, []int{2, 4}, 1))
	fmt.Println(work(10, []int{2, 4, 7}, 1))
	fmt.Println(work(5, []int{2, 4}, 2))
}

func work(all int, dead []int, k int) int {
	m := make(map[int]struct{})
	for _, v := range dead {
		m[v] = struct{}{}
	}
	var curDead, maxLen int
	l, r := 1, 1
	for r <= all {
		if _, ok := m[r]; ok {
			curDead++
			for curDead > k {
				if _, ok := m[l]; ok {
					curDead--
				}
				l++
			}
		}
		r++
		if r-l > maxLen {
			maxLen = r - l
		}
	}
	return maxLen
}
