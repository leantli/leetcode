package main

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 104. 二叉树的最大深度

// 最好的写法，自底向上递归，从叶子节点开始计算层数
// 遍历到最底，为 nil，此时返回 0
// 其他的直接返回其左右子节点的最大深度 + 1(其本身)
// 整个过程可以看成是后序遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// // 自顶向下递归，逐层往下计算深度，到叶子节点则更新最大深度
// // 整个过程可以看成是前序遍历
// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	var depth int
// 	var travel func(root *TreeNode, cur int)
// 	travel = func(root *TreeNode, cur int) {
// 		if root == nil {
// 			return
// 		}
// 		if root.Left == nil && root.Right == nil {
// 			if depth < cur {
// 				depth = cur
// 			}
// 		}
// 		travel(root.Left, cur+1)
// 		travel(root.Right, cur+1)
// 	}
// 	travel(root, 1)
// 	return depth
// }

// // 任意一种遍历，记录最深的即可？
// func maxDepth(root *TreeNode) int {
// 	var depth int
// 	var cur int
// 	var travel func(root *TreeNode)
// 	travel = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		cur++
// 		if cur > depth {
// 			depth = cur
// 		}
// 		travel(root.Left)
// 		travel(root.Right)
// 		cur--

// 	}
// 	travel(root)
// 	return depth
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
