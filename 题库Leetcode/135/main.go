package main

// https://leetcode.cn/problems/candy/
// 135. 分发糖果

// 二刷，从左、右各遍历一次，分别找到各个女孩应该收到的糖果数量
// 最终取 max 即可
func candy(ratings []int) int {
	left := make([]int, len(ratings))
	left[0] = 1
	for i := 1; i < len(ratings); i++ {
		left[i] = 1
		if ratings[i] > ratings[i-1] {
			left[i] += left[i-1]
		}
	}
	right := 1
	cnt := max(right, left[len(ratings)-1])
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right += 1
		} else {
			right = 1
		}
		cnt += max(right, left[i])
	}
	return cnt
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// // 贪心，我们就按题目要求来，每个孩子至少给一颗糖果
// // 如果该孩子评分比旁边的高，就给他，比旁边的多一颗糖果即可
// // 因此从左、从右各遍历一次即可，选取每个孩子可能拿到的最大糖果数量即可
// func candy(ratings []int) int {
// 	n := len(ratings)
// 	left := make([]int, n)
// 	for i := range left {
// 		if i > 0 && ratings[i] > ratings[i-1] {
// 			left[i] = left[i-1] + 1
// 		} else {
// 			left[i] = 1
// 		}
// 	}
// 	var right, res int
// 	for i := n - 1; i >= 0; i-- {
// 		if i < n-1 && ratings[i] > ratings[i+1] {
// 			right++
// 		} else {
// 			right = 1
// 		}
// 		res += max(left[i], right)
// 	}
// 	return res
// }
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
