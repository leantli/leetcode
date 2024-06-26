package main

// https://leetcode.cn/problems/combinations/
// 77. 组合

// // 返回 1~n 中所有可能的 k 个数字的组合
// func combine(n int, k int) [][]int {
// 	res := make([][]int, 0)
// 	cur := make([]int, 0)
// 	var dfs func(idx int)
// 	dfs = func(idx int) {
// 		if len(cur) == k {
// 			res = append(res, append([]int{}, cur...))
// 			return
// 		}
// 		for i := idx; i <= n; i++ {
// 			cur = append(cur, i)
//          // i 已经被选取，所以后续从 i+1 开始选取
// 			dfs(i + 1)
// 			cur = cur[:len(cur)-1]
// 		}
// 	}
// 	dfs(1)
// 	return res
// }

// 剪枝
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	var dfs func(idx int)
	dfs = func(idx int) {
		if len(cur) == k {
			res = append(res, append([]int{}, cur...))
			return
		}
		// 可选取的 candidate [idx,n]
		for i := idx; i <= n; i++ {
			// 当 cur 长度 + (n-i+1) 数量的元素，小于 k 时，就不用再考虑 dfs 了
			// 因为把后面的元素加在 cur 上，也达不到数量 k
			if len(cur)+n-i+1 < k {
				break
			}
			cur = append(cur, i)
			dfs(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(1)
	return res
}

// // 另一种 dfs 写法，只判断当前的数取或不取，两种 dfs 路线
// func combine(n int, k int) [][]int {
// 	res := make([][]int, 0)
// 	cur := make([]int, 0)
// 	var dfs func(idx int)
// 	dfs = func(idx int) {
// 		// 剪枝，当 cur 当前的长度加上后续 [idx,n] 的所有元素也达不到 k 的长度时，直接放弃
// 		if len(cur)+(n-idx) < k {
// 			return
// 		}
// 		// 满足条件时，放入结果集中
// 		if len(cur) == k {
// 			res = append(res, append([]int{}, cur...))
// 			return
// 		}
// 		// 跳过当前整数不取
// 		dfs(idx + 1)
// 		// 不跳过则取 idx 下标
// 		cur = append(cur, idx+1)
// 		dfs(idx + 1)
// 		cur = cur[:len(cur)-1]
// 	}
// 	dfs(0)
// 	return res
// }
