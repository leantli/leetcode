package main

// https://leetcode.cn/problems/binary-tree-preorder-traversal/
// 144. 二叉树的前序遍历

// 模拟栈 实现
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	// 当节点不为空且栈中仍有节点，就继续遍历
	for root != nil || len(stack) != 0 {
		// root 不为空时，说明当前root的中左还没遍历完，优先中左
		if root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			// root 为 nil ，说明一开始的 root 的中左序列都遍历完了
			// 现在我们看看这个中左序列上，是否有哪个节点有右子结点
			// 有的话，就赋值给 root 并继续中左遍历
			root = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
	return res
}

// 另一种 模拟栈 的方式实现前序遍历, 根左右，较易理解，但和中序后序模拟栈写法不相通
// func preorderTraversal(root *TreeNode) []int {
// 	res := make([]int, 0)
// 	if root == nil {
// 		return res
// 	}
// 	stack := make([]*TreeNode, 0)
// 	stack = append(stack, root)
// 	// 栈首先加入 root 节点
// 	// 先打印 根结点，由于后续是先左后右
// 	// 因此栈中先入右节点，再入左节点
// 	for len(stack) != 0 {
// 		node := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		res = append(res, node.Val)
// 		if node.Right != nil {
// 			stack = append(stack, node.Right)
// 		}
// 		if node.Left != nil {
// 			stack = append(stack, node.Left)
// 		}
// 	}
// 	return res
// }

// // 基础，理解前序遍历, 根左右
// func preorderTraversal(root *TreeNode) []int {
// 	res := make([]int, 0)
// 	var preorederTravel func(root *TreeNode)
// 	preorederTravel = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		res = append(res, root.Val)
// 		preorederTravel(root.Left)
// 		preorederTravel(root.Right)
// 	}
// 	preorederTravel(root)
// 	return res
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
