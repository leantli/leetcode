package main

// https://leetcode.cn/problems/binary-tree-right-side-view/description/
// 199. 二叉树的右视图

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 如何使用到递归？显然应该是先序遍历，同时右子结点的遍历优先于左子节点
// 当当前的深度等于字符串长度时，说明当前遍历到的节点就是树每一层右侧起第一个节点
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(res) {
			res = append(res, node.Val)
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}
	dfs(root, 0)
	return res
}

// // 层序遍历，每层都从右节点开始存入，这样能够保证，每一层的第一个节点，就是我们要右视图的值
// func rightSideView(root *TreeNode) []int {
// 	res := make([]int, 0)
// 	if root == nil {
// 		return res
// 	}
// 	queue := []*TreeNode{root}
// 	for len(queue) > 0 {
// 		n := len(queue)
// 		res = append(res, queue[0].Val)
// 		for i := 0; i < n; i++ {
// 			if queue[i].Right != nil {
// 				queue = append(queue, queue[i].Right)
// 			}
// 			if queue[i].Left != nil {
// 				queue = append(queue, queue[i].Left)
// 			}
// 		}
// 		queue = queue[n:]
// 	}
// 	return res
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
