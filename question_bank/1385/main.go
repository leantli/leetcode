package main

import "sort"

// https://leetcode.cn/problems/find-the-distance-value-between-two-arrays/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 1385. 两个数组间的距离值

// 写完模拟之后，其实我们能发现有很多判断是不必要的，我们只需要判断 arr2 中是否有数字在 num1 +- 2 范围内即可，有则不符合距离要求，无需加入距离值
// arr2 先一次排序 (mlogm), 在根据 n 中每个数，对 arr2 二分找 num-d ～ num+d 范围内的值， (nlogm)
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})
	var ans int
	for _, num := range arr1 {
		if isOutOfRule(arr2, num-d, num+d) {
			continue
		}
		ans++
	}
	return ans
}

// 判断是否不符合距离要求，是则返回 true
func isOutOfRule(arr2 []int, low, high int) bool {
	l, r := 0, len(arr2)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr2[mid] >= low && arr2[mid] <= high {
			return true
		} else if arr2[mid] < low {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

// 模拟？
// func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
// 	var ans int
// 	n := len(arr2)
// 	for _, num1 := range arr1 {
// 		count := 0
// 		for _, num2 := range arr2 {
// 			if abs(num1, num2) > d {
// 				count++
// 			}
// 		}
// 		if count == n {
// 			ans++
// 		}
// 	}
// 	return ans
// }

func abs(a, b int) int {
	temp := a - b
	if temp > 0 {
		return temp
	}
	return -temp
}
