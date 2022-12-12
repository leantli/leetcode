package main

// https://leetcode.cn/problems/er-cha-shu-de-shen-du-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 55-1. 二叉树的深度

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 官解更简洁，打扰了
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

// // 好像没啥好分析，就树的遍历，然后 depth 记一下层数以及最大深度
// func maxDepth(root *TreeNode) int {
// 	var res int
// 	var dfs func(root *TreeNode, depth int)
// 	dfs = func(root *TreeNode, depth int) {
// 		if root != nil {
// 			depth++
// 			if res < depth {
// 				res = depth
// 			}
// 			dfs(root.Left, depth)
// 			dfs(root.Right, depth)
// 		}
// 	}
// 	dfs(root, 0)
// 	return res
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
