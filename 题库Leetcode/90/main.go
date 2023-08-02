package main

import "sort"

// https://leetcode.cn/problems/subsets-ii/
// 90.子集II

// nums 中可能存在重复的数，选取的子集不能重复出现
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	sort.Ints(nums)
	var dfs func(startIdx int)
	dfs = func(startIdx int) {
		res = append(res, append([]int{}, cur...))
		if startIdx == len(nums) {
			return
		}
		for i := startIdx; i < len(nums); i++ {
			// 去重，排序后相邻数一致，只用选取一次即可
			if i > startIdx && nums[i] == nums[i-1] {
				continue
			}
			cur = append(cur, nums[i])
			dfs(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0)
	return res
}
