package main

// https://leetcode.cn/problems/invert-binary-tree/
// 226 翻转二叉树

// 遍历每个节点，并将每个节点下的左右子节点调换位置即可
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
