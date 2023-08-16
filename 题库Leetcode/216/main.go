package main

// 216. 组合总和 III
// https://leetcode.cn/problems/combination-sum-iii/

// 找到和为 n 的 k 个数的组合，每个数只能用一次，从 1-9 内挑选组合
func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	var sum int
	var dfs func(idx int)
	dfs = func(idx int) {
		if sum == n && len(cur) == k {
			res = append(res, append([]int{}, cur...))
			return
		}
		if sum >= n {
			return
		}
		for i := idx; i <= 9; i++ {
			// 剪枝
			// 如果当前的 cur 长度，加上后面的所有元素，数量都达不到 k，就直接 break
			if 10-i+len(cur) < k {
				break
			}
			cur = append(cur, i)
			sum += i
			dfs(i + 1)
			cur = cur[:len(cur)-1]
			sum -= i
		}
	}
	dfs(1)
	return res
}
