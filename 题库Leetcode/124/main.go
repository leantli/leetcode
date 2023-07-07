package main

import "math"

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/
// 124. 二叉树中的最大路径和

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt
	var getMax func(root *TreeNode) int
	getMax = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 自底向上，获取左右子树下最大值的路径，如果得到的路径和小于 0，直接抛弃
		leftMax := max(getMax(root.Left), 0)
		rightMax := max(getMax(root.Right), 0)
		// 当前节点连通左右路径的最大和
		curMax := root.Val + leftMax + rightMax
		// 更新最大值
		maxSum = max(maxSum, curMax)
		// 左右路径只能选一条和当前节点一起返回，当然选最大的
		return root.Val + max(leftMax, rightMax)
	}
	getMax(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
