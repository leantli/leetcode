package main

// https://leetcode.cn/problems/binary-tree-inorder-traversal/
// 94. 二叉树的中序遍历

// 模拟栈实现 左根右
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := []*TreeNode{}
	for root != nil || len(stack) != 0 {
		// root 不为空时，说明 root 的左中还未遍历完，优先左中
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			// 取出左(根)节点，加入res中
			// 将其右子结点赋值给 root，不管新的 root 是 nil 还是有值
			// 都依然优先左中，后右的顺序遍历取值
			res = append(res, stack[len(stack)-1].Val)
			root = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
	return res
}

// // 二叉树中序遍历 左根右
// func inorderTraversal(root *TreeNode) []int {
// 	res := make([]int, 0)
// 	var inorderTravel func(root *TreeNode)
// 	inorderTravel = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		inorderTravel(root.Left)
// 		res = append(res, root.Val)
// 		inorderTravel(root.Right)
// 	}
// 	inorderTravel(root)
// 	return res
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
