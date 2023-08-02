package main

import "sort"

// https://leetcode.cn/problems/permutations-ii/
// 47.全排列 II

// 可能包含重复数字，全排列
func permuteUnique(nums []int) [][]int {
	// 重复数，相邻
	sort.Ints(nums)
	res := make([][]int, 0)
	cur := make([]int, 0)
	used := make([]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(cur) == len(nums) {
			res = append(res, append([]int{}, cur...))
			return
		}
		for i := 0; i < len(nums); i++ {
			// 当 nums[i] == nums[i-1] 时，并且 i-1 为未使用过，说明这层已经被选取过了
			// 选过之后，会将 used 置为 false，相同的数，当前位置就不必再重复选取
			if i > 0 && nums[i] == nums[i-1] && used[i-1] {
				continue
			}
			// 只能使用还未被使用过的数字
			if !used[i] {
				used[i] = true
				cur = append(cur, nums[i])
				// 选取下一个位置
				dfs()
				cur = cur[:len(cur)-1]
				used[i] = false
			}
		}
	}
	dfs()
	return res
}
