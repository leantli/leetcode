package main

// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/
// 68-2. 二叉树的最近公共祖先

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 怎么找到最近公共祖先？ root 左右子节点都找到了 p 和 q ？
// 递归地去找 p 和 q
// 找到了就返回
// 在上层 root 判断，left 和 right 是否有找到 p 和 q
// 如果两边都找到了，显然 root 就是最近公共祖先
// 如果只有一边找到了，就先只返回一边
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
