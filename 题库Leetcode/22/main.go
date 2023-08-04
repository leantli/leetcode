package main

// https://leetcode.cn/problems/generate-parentheses/
// 22. 括号生成

// dfs，n 表示左括号和右括号的个数，dfs 枚举每种情况，但是我们需要注意剪枝，右括号的使用数量必须时刻小于等于左括号的使用数量，否则括号组合无效
func generateParenthesis(n int) []string {
	var dfs func(l, r int)
	var cur string
	res := make([]string, 0)
	dfs = func(l, r int) {
		// 当右括号的使用数量到达 n，说明已经完成组合，加入结果中
		if r == n {
			res = append(res, cur)
			return
		}
		// 当左括号数量小于 n 时，尝试增加左括号
		if l < n {
			cur = cur + "("
			dfs(l+1, r)
			cur = cur[:len(cur)-1]
		}
		// 当右括号的数量小于左括号时，尝试增加右括号
		if r < l {
			cur = cur + ")"
			dfs(l, r+1)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0, 0)
	return res
}
