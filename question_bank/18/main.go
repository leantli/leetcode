package main

import "sort"

// https://leetcode.cn/problems/4sum/
// 18. 四数之和

// 和三数之和的思路基本一样，由于无需对应的位置只需要值，因此可以先排个序
// 其次就是第一个值就比 target 大，肯定可以直接 break 了
// 已经循环时不和前一个数相同，避免重复统计
// 当找到满足条件的四个数时，后两个数需要继续步进排重
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for a := 0; a < len(nums)-3; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b < len(nums)-2; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			c, d := b+1, len(nums)-1
			for c < d {
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum == target {
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
					c++
					d--
					for c < d && nums[c] == nums[c-1] {
						c++
					}
					continue
				}
				if sum > target {
					d--
				} else {
					c++
				}
			}
		}
	}
	return res
}
