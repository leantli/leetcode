package main

// https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 55-2. 平衡二叉树

// 任意节点的左右子树深度都不超过 1，否则返回 false
// 先写出获取节点深度的函数，然后判断 root 的左右子树是否满足条件，不满足则返回 false
// 再递归调用当前函数，对左子节点和右子结点进行操作
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(getDeepth(root.Left), getDeepth(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getDeepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getDeepth(root.Left), getDeepth(root.Right)) + 1
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

// // better solution 在 height 取层高的同时判断左右子树的深度情况
// func isBalanced(root *TreeNode) bool {
// 	return height(root) >= 0
// }

// func height(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	leftHeight := height(root.Left)
// 	rightHeight := height(root.Right)
// 	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
// 		return -1
// 	}
// 	return max(leftHeight, rightHeight) + 1
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -1 * x
// 	}
// 	return x
// }

// 我的
// func isBalanced(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	if res := isBalanced(root.Left) && isBalanced(root.Right); !res {
// 		return false
// 	}
// 	diff := maxDepth(root.Left) - maxDepth(root.Right)
// 	if diff > 1 || diff < -1 {
// 		return false
// 	}
// 	return true
// }

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
