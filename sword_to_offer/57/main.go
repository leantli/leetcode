package main

// https://leetcode.cn/problems/he-wei-sde-liang-ge-shu-zi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 57. 和为s的两个数字

// 两数和 map 也行，但是这个递增

func twoSum(nums []int, target int) []int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		sum := nums[l] + nums[r]
		if sum == target {
			return []int{nums[l], nums[r]}
		}
		if sum > target {
			r--
		}
		if sum < target {
			l++
		}
	}
	return []int{}
}
