package main

import "sort"

// https://leetcode.cn/problems/zi-fu-chuan-de-pai-lie-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 38. 字符串的排列

// 等同于 leetcode 47. 全排列 II 思路，去重有三种，基于set对叶子节点去重；基于sort对树层或树枝去重

// 补充对树层去重，先排序，后在遍历 bs[i] 时，查看 used[i] 是否被用过 或者 bs[i] == bs[i-1] && used[i-1] = false
// 此时可以跳过 bs[i] 的选取，因为之前已经发生过 bs[i] 作为这个坑的元素的 dfs 了, 去重操作
// 具体分析见 https://www.programmercarl.com/0047.%E5%85%A8%E6%8E%92%E5%88%97II.html#c-%E4%BB%A3%E7%A0%81
func permutation(s string) []string {
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
	used := make([]bool, len(bs))
	res := make([]string, 0)
	var dfs func(cur string)
	dfs = func(cur string) {
		if len(cur) == len(s) {
			res = append(res, cur)
			return
		}
		for i := range bs {
			if used[i] || (i > 0 && bs[i] == bs[i-1] && !used[i-1]) {
				continue
			}
			used[i] = true
			dfs(cur + string(bs[i]))
			used[i] = false
		}
	}
	dfs("")
	return res
}

// // 填坑递归 + set
// func permutation(s string) []string {
// 	set := make(map[string]struct{})
// 	used := make([]bool, len(s))
// 	var dfs func(cur string)
// 	dfs = func(cur string) {
// 		// 当长度够了直接加进set，终止递归
// 		if len(cur) == len(s) {
// 			set[cur] = struct{}{}
// 			return
// 		}
// 		// 遍历 s[i]，看看有没有被用过，没用过的配合 cur 继续 dfs
// 		for i := range s {
// 			if !used[i] {
// 				used[i] = true
// 				dfs(cur + string(s[i]))
// 				used[i] = false
// 			}
// 		}
// 	}
// 	dfs("")
// 	res := make([]string, 0, len(set))
// 	for k, _ := range set {
// 		res = append(res, k)
// 	}
// 	return res
// }
