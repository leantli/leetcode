package main

// https://leetcode.cn/problems/largest-rectangle-in-histogram/
// 84. 柱状图中最大的矩形

// 还是单调栈，单调递减还是递增？如何考虑？我们想到，每一列左右都要扩的话，左右都必须比当前的大
// 假设是单调递减栈，从栈底到栈顶单调递减，此时遇到比栈顶大的下标，弹出栈顶的下标作为mid，此时 mid 左右确实都可以比当前大，但是栈中的所有下标都可以作为边界，显然是不行的
// 假设是单调递增栈，从栈底到栈顶单调递增，此时遇到比栈顶小的下标，弹出栈顶的下标作为mid，此时 mid 左右下标都比当前小，因此可以计算 mid * (r-l-1)
// 比如说 1 5 6 2，栈 [1 5 6]，此时遇到 2，弹出 6，面积为 6*(3-1-1)，此时栈顶下标为 1，指向 5，仍然比 2 大，弹出 5，此时面积为 5*(3-0-1) = 10
// 同时，为了保证边界值都被处理到，我们可以在 heights 左右各自增加一个零值，保证边界也能正常处理
func largestRectangleArea(heights []int) int {
	// 在 heights 首尾增加 0，保证边界能被正常处理
	heights = append([]int{0}, append(heights, 0)...)
	stack := make([]int, 0)
	var res int
	for r, rightHeight := range heights {
		// 当栈中有数并且栈顶的下标的高度大于当前 rightHeight，则进行面积计算
		for len(stack) > 0 && rightHeight < heights[stack[len(stack)-1]] {
			midHeight := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				res = max(res, midHeight*(r-stack[len(stack)-1]-1))
			}
		}
		stack = append(stack, r)
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// // 最朴素的思路，以每一列的值为矩形的高，不断向左右拓展，直到遇到不能拓展的边界(高度小于当前列值的)
// // 此时计算对应的面积，时间复杂度为 O(n^2)，我们可以先计算出每一列左右第一个小于该列的下标位置
// // 方便后续计算，本题应该会超时，否则不可能为困难题
// func largestRectangleArea(heights []int) int {
// 	n := len(heights)
// 	L, R := make([]int, n), make([]int, n)
// 	// L[i] 表示 i 下标高度左侧第一个小于 height[i] 的下标位置
// 	for i := 0; i < n; i++ {
// 		l := i - 1
// 		for l >= 0 && heights[l] >= heights[i] {
// 			l--
// 		}
// 		L[i] = l
// 	}
// 	// R[i] 表示 i 下标高度右侧第一个小于 height[i] 的下标位置
// 	for i := n - 1; i >= 0; i-- {
// 		r := i + 1
// 		for r <= n-1 && heights[r] >= heights[i] {
// 			r++
// 		}
// 		R[i] = r
// 	}
// 	var res int
// 	for i, height := range heights {
// 		res = max(res, height*(R[i]-L[i]-1))
// 	}
// 	return res
// }
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
