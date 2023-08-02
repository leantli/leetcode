package main

// https://leetcode.cn/problems/permutations/
// 46. 全排列

// 不含重复数字， 全排列
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	used := make([]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(cur) == len(nums) {
			res = append(res, append([]int{}, cur...))
			return
		}
		// 当前层，每个位置的candidate，只要没被选过都可以使用
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				used[i] = true
				cur = append(cur, nums[i])
				dfs()
				cur = cur[:len(cur)-1]
				used[i] = false
			}
		}
	}
	dfs()
	return res
}
