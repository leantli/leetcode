package main

// https://leetcode.cn/problems/binary-tree-postorder-traversal/
// 145. 二叉树的后序遍历

// 模拟栈实现 左右中
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := []*TreeNode{}
	// 增加一个节点指针，指向上次访问的节点，用于判断，这次到根结点
	// 是从左子节点回溯回来的，还是从右子结点回溯回来的
	var pre *TreeNode
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			// 该节点存在右子结点并且上一个访问的不是右子结点
			// 证明其实从左子节点回溯过来的，因此接下来要访问右子结点
			if root.Right != nil && pre != root.Right {
				root = root.Right
			} else {
				// 该节点不存在右子结点或已经访问过右子结点
				// 此时访问根结点
				res = append(res, root.Val)
				pre = root
				root = nil
				stack = stack[:len(stack)-1]
			}
		}
	}
	return res
}

// // 递归实现后序遍历 左右根
// func postorderTraversal(root *TreeNode) []int {
// 	res := make([]int, 0)
// 	var postorderTravel func(root *TreeNode)
// 	postorderTravel = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		postorderTravel(root.Left)
// 		postorderTravel(root.Right)
// 		res = append(res, root.Val)
// 	}
// 	postorderTravel(root)
// 	return res
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
