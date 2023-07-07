package main

// https://leetcode.cn/problems/candy/
// 135. 分发糖果

// 贪心，我们就按题目要求来，每个孩子至少给一颗糖果
// 如果该孩子评分比旁边的高，就给他，比旁边的多一颗糖果即可
// 因此从左、从右各遍历一次即可，选取每个孩子可能拿到的最大糖果数量即可
func candy(ratings []int) int {
	n := len(ratings)
	left := make([]int, n)
	for i := range left {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	var right, res int
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		res += max(left[i], right)
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
