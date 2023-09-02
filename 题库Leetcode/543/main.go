package main

// https://leetcode.cn/problems/diameter-of-binary-tree/description/
// 543. 二叉树的直径

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 直观感觉，获取每个节点左右子树的深度，将左右子树的深度相加，就是经过当前节点的最长路径
// 对每个节点都操作一遍，获取最大的路径节点有多少个，长度等于节点数减一即可
func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l, r := dfs(root.Left), dfs(root.Right)
		if l+r > res {
			res = l + r
		}
		return max(l, r) + 1
	}
	dfs(root)
	return res
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
