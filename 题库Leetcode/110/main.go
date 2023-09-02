package main

// https://leetcode.cn/problems/balanced-binary-tree/
// 110. 平衡二叉树

// 二刷
// 判断是否是平衡二叉树，显然需要自底向上进行递归判断
// 递归函数需要返回递归节点的深度，便于父节点判断其本身是否为平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l, r := getDepth(root.Left), getDepth(root.Right)
	if abs(l, r) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

// 获取左右子树的深度，选取更大的深度+1作为当前节点的深度
func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(getDepth(node.Left), getDepth(node.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a, b int) int {
	sum := a - b
	if sum < 0 {
		return -sum
	}
	return sum
}

// // 任意节点的左右子树深度都不超过 1，否则返回 false
// // 先写出获取节点深度的函数，然后判断 root 的左右子树是否满足条件，不满足则返回 false
// // 再递归调用当前函数，对左子节点和右子结点进行操作
// func isBalanced(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	if abs(getDeepth(root.Left), getDeepth(root.Right)) > 1 {
// 		return false
// 	}
// 	return isBalanced(root.Left) && isBalanced(root.Right)
// }

// func getDeepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return max(getDeepth(root.Left), getDeepth(root.Right)) + 1
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func abs(a, b int) int {
// 	sum := a - b
// 	if sum < 0 {
// 		return -sum
// 	}
// 	return sum
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
