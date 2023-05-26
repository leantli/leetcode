package main

// https://leetcode.cn/problems/jian-sheng-zi-ii-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 14-2. 剪绳子 II

// 这道题我乍看之下感觉和 14-1 没区别，再仔细观察才发现，原理是多了对大数的处理
// 整体思路和 14-1 差不多， dfs 超时，dp n^2, 数学 n
// 其中 dp 对大数的处理不能在状态转移过程中处理，只能等到最后处理，否则中间各种 max(a,b) 比较时可能会误判，但结果很可能超过 int64，因此不得不使用 big.int 库对大数做处理
// 但数学思路时，可以循环取余，因为不存在其他判断

// 3m+k = n
func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}
	m, k := n/3, n%3
	res, mod := 1, 1000000007
	for i := 0; i < m-1; i++ {
		res = (res * 3) % mod
	}
	switch k {
	case 1:
		return (res * 4) % mod
	case 2:
		return (res * 6) % mod
	}
	return (res * 3) % mod
}
