package main

import "sort"

// https://leetcode.cn/problems/sum-of-mutated-array-closest-to-target/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=bv79h7h
// 1300. 转变数组后最接近目标值的数组和

// 二分，模拟，前缀和 思路
// 要找到一个 value，该 value 会使 arr 数组中比它大的都缩小为它
// 最终 arr 数组之和越接近 target 越好，返回最好的 value，同样接近取值更小的 value
// 这里最好再用一下前缀和
// var sum int

// func findBestValue(arr []int, target int) int {
// 	var maxNum int
// 	sum = 0
// 	for _, v := range arr {
// 		sum += v
// 		maxNum = max(maxNum, v)
// 	}
// 	l, r := 0, maxNum
// 	for l+1 != r {
// 		mid := l + (r-l)/2
// 		// fmt.Println(mid)
// 		if calSum(arr, mid) < target {
// 			l = mid
// 		} else {
// 			r = mid
// 		}
// 	}
// 	// 此时 l 为最大的 导致数组和 小于 target 的 value
// 	// r 为最小的 导致数组和 大于等于 target 的 value
// 	if abs(calSum(arr, l), target) <= abs(calSum(arr, r), target) {
// 		return l
// 	}
// 	return r
// }

// func abs(a, b int) int {
// 	res := a - b
// 	if res < 0 {
// 		return -res
// 	}
// 	return res
// }

// func calSum(arr []int, m int) int {
// 	res := sum
// 	for _, v := range arr {
// 		// 由于我们已经有前缀和 sum，接下来只需要把相应超出的部分减去即可
// 		if v > m {
// 			res -= v - m
// 		}
// 	}
// 	return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// 感觉更像数学思路？
func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	for _, v := range arr {
		curAvg := float64(target) / float64(n)
		if float64(v) < curAvg {
			target -= v
			n--
			continue
		}
		if curAvg-float64(int(curAvg)) <= 0.5 {
			return int(curAvg)
		} else {
			return int(curAvg) + 1
		}
	}
	return arr[len(arr)-1]
}
