package main

import "sort"

// https://leetcode.cn/problems/boats-to-save-people/
// 881. 救生艇

// 在下面的基础上简单优化了一下结构
func numRescueBoats(people []int, limit int) int {
	var res int
	sort.Ints(people)
	l, r := 0, len(people)-1
	for l <= r {
		if people[l]+people[r] <= limit {
			// 匹配成功
			l++
		}
		// 不管是否匹配成功，r都会左移，船数量也会增加
		res++
		r--
	}
	return res
}

// // 显然是要不断去匹配最轻和最重的，重的匹配不上的就只能单独
// func numRescueBoats(people []int, limit int) int {
// 	var res int
// 	sort.Ints(people)
// 	l, r := 0, len(people)-1
// 	for l <= r {
// 		sum := people[l] + people[r]
// 		if sum <= limit {
// 			// 匹配成功
// 			res++
// 			l++
// 			r--
// 		} else {
// 			// 没匹配上，重的单独一艘
// 			res++
// 			r--
// 		}
// 	}
// 	return res
// }
