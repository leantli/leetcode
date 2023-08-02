package main

// https://leetcode.cn/problems/combination-sum/
// 39. 组合总和

// 无重复元素，可被无限制重复选取，数字和为 target，从 candidates 中选取
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	var sum int
	var dfs func(idx int)
	dfs = func(idx int) {
		if sum > target {
			return
		}
		if sum == target {
			res = append(res, append([]int{}, cur...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			cur = append(cur, candidates[i])
			sum += candidates[i]
			dfs(i)
			sum -= candidates[i]
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0)
	return res
}
