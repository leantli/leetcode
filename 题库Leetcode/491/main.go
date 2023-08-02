package main

// https://leetcode.cn/problems/non-decreasing-subsequences/
// 491.递增子序列

// 返回所有递增子序列，子序列至少两个元素，数组中存在重复元素(子序列无法排序去重)
func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	var dfs func(startIdx int)
	dfs = func(startIdx int) {
		if len(cur) >= 2 {
			res = append(res, append([]int{}, cur...))
		}
		if startIdx == len(nums) {
			return
		}
		// 因此我们需要看，在同一个位置的选取上，该数是否已经遇到过
		// 遇到过则已经输出过相关的递增子序列了，不需要再重复选取
		met := make(map[int]struct{})
		for i := startIdx; i < len(nums); i++ {
			if _, ok := met[nums[i]]; ok {
				continue
			}
			if len(cur) == 0 || nums[i] >= cur[len(cur)-1] {
				met[nums[i]] = struct{}{}
				cur = append(cur, nums[i])
				dfs(i + 1)
				cur = cur[:len(cur)-1]
			}
		}
	}
	dfs(0)
	return res
}
