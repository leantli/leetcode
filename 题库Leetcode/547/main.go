package main

// https://leetcode.cn/problems/number-of-provinces/
// 547. 省份数量

// 并查集
// set[i] 表示 i 的父节点的index，默认情况下 set[i] = i
// find(x int)，查找 x 的父节点，并且路径压缩，set[x] = find(set[x])，递归过程中同时进行赋值
// union(from, to int)，统一 from 和 to 的父节点
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	set := make([]int, n)
	for i := range set {
		set[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if set[x] != x {
			set[x] = find(set[x])
		}
		return set[x]
	}
	union := func(from, to int) {
		set[find(from)] = find(to)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}
	var res int
	for i := range set {
		if i == set[i] {
			res++
		}
	}
	return res
}
