package main

import "sort"

// https://leetcode.cn/problems/3sum-closest/
// 16. 最接近的三数之和

// 最接近的三数之和，只需返回最终的和即可
// 我们可以先排序，但二分还是不太行，因为二分的话，我们得先确定 a 和 b 的位置
// 最后再对 c 进行 logn 的二分，但是 a 和 b 位置的确定不就已经 n^2 了？
// 所以还是常规双指针，然后去取 abs 最小的即可
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	diff := 20001
	var res int
	for a := 0; a < len(nums)-2; a++ {
		b, c := a+1, len(nums)-1
		for b < c {
			sum := nums[a] + nums[b] + nums[c]
			if sum <= target {
				b++
			} else {
				c--
			}
			curDiff := abs(sum, target)
			if curDiff < diff {
				res = sum
				diff = curDiff
			}
		}
	}
	return res
}

func abs(a, b int) int {
	res := a - b
	if res < 0 {
		return -res
	}
	return res
}
