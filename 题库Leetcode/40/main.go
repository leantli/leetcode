package main

import (
	"sort"
)

// https://leetcode.cn/problems/combination-sum-ii/
// 40.组合总和II

// 从 candidates 中找出一些组合，这些组合的 sum = target，每个数只能用一次
// candidates 中存在重复的数字，可以先 sort 一遍，在后续的选取过程中，进行过滤
func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	var sum int
	sort.Ints(candidates)
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
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			cur = append(cur, candidates[i])
			sum += candidates[i]
			// 选取下一个还未选取过的
			dfs(i + 1)
			sum -= candidates[i]
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0)
	return res
}
