package main

import "sort"

// https://leetcode.cn/problems/check-if-n-and-its-double-exist/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 1346. 检查整数及其两倍数是否存在

// 排序+二分 == nlogn+nlogn = 2nlogn
func checkIfExist(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	n := len(arr)
	for i := 0; i < n; i++ {
		// 这里对 l 和 r 的取值范围要简单处理一下
		var l, r int
		if arr[i] >= 0 {
			l, r = i+1, n-1
		} else {
			l, r = 0, i-1
		}
		for l <= r {
			mid := l + (r-l)/2
			temp := arr[i] * 2
			if arr[mid] == temp {
				return true
			} else if arr[mid] > temp {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return false
}

// // 哈希感觉更方便与快速，时间复杂度为 O(n)
// // 但是还是要注意，对 0 要特殊处理一下。。。
// func checkIfExist(arr []int) bool {
// 	m := make(map[int]int)
// 	for _, num := range arr {
// 		m[num]++
// 	}
// 	for k := range m {
// 		res := m[k*2]
// 		if k == 0 {
// 			if res >= 2 {
// 				return true
// 			}
// 		} else {
// 			if res >= 1 {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
