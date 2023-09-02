package main

// https://leetcode.cn/problems/validate-binary-search-tree/
// 98. 验证二叉搜索树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 我们不能只验证单个节点的左子节点小于当前节点，右子结点大于当前节点
// 因为可能存在左子节点的右孙子节点比当前节点大
// 因此我们不在递归中去比较当前节点和左右子节点，而要合理利用二叉搜索树的性质
// 二叉搜索树的中序遍历即是单调非递减的，因此我们可以使用一个局部变量去记录之前遇到的最后一个数值
// 每次比较时，当前数比之前最后一个数值大于等于即可，比较成功则替换并继续中序遍历
func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	res := true
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil || !res {
			return
		}
		inorder(root.Left)
		// 当 pre 为空 或 当前节点的值大于前置节点值，则正常遍历
		if pre == nil || root.Val > pre.Val {
			pre = root
		} else {
			res = false
		}
		inorder(root.Right)
	}
	inorder(root)
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
