package main

// https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 27. 二叉树的镜像

// 遍历每个节点，并将每个节点下的左右子节点调换位置即可
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
