package main

// https://leetcode.cn/problems/subsets/
// 78. 子集

// 返回所有可能的子集，枚举所有可能的情况
// 和组合非常像，但组合只需要返回最终的叶子节点，而现在的需求是获取搜索过程上的所有节点值
// 并且注意的是，[] 空集合也算在内
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	var dfs func(startIdx int)
	dfs = func(startIdx int) {
		res = append(res, append([]int{}, cur...))
		if startIdx == len(nums) {
			return
		}
		for i := startIdx; i < len(nums); i++ {
			cur = append(cur, nums[i])
			dfs(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0)
	return res
}
