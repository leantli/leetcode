package main

// https://leetcode.cn/problems/count-complete-tree-nodes/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 222. 完全二叉树的节点个数

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 活用完全二叉树的性质？
// 当两个层数相等时，左子树肯定是满的，只要根据层数即可直接计算左子树的全部节点数
// 当两个层数不等时，右子树肯定是满的，根据层数算右子树的全部节点数量
// 并且完全二叉树的子树也符合完全二叉树的性质，因此也可以基于该思路递归去计算
// 总体时间复杂度应该是 logn*logn
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lD := calDepth(root.Left)
	rD := calDepth(root.Right)
	if lD == rD {
		// return countNodes(root.Right) + (1 << lD) - 1 + 1
		return countNodes(root.Right) + (1 << lD)
	}
	// return countNodes(root.Left) + (1 << rD) - 1 + 1
	return countNodes(root.Left) + (1 << rD)
}

// 计算某个节点的深度
func calDepth(root *TreeNode) int {
	var depth int
	for root != nil {
		depth++
		root = root.Left
	}
	return depth
}

// // 最差的全遍历 O(n)
// func countNodes(root *TreeNode) int {
// 	var cnt int
// 	var fun func(*TreeNode)
// 	fun = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		cnt++
// 		fun(root.Left)
// 		fun(root.Right)
// 	}
// 	fun(root)
// 	return cnt
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
