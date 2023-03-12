package main

// https://leetcode.cn/problems/symmetric-tree/
// 101. 对称二叉树

// 递归的方式--自顶向下
// 这边的递归入参必须是两个节点，这样才能传入每个节点的左右子节点
// 保证左右子树中，最右侧和最左侧的节点对称
// 即题目要求对称，显然需要左子树中的节点和右子树中的节点比较，显然需要传两个参数
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 比较左右节点是否一致
	var compare func(left, right *TreeNode) bool
	compare = func(left, right *TreeNode) bool {
		// 左右节点都为空，一致
		if left == nil && right == nil {
			return true
		}
		// 左右节点非空且值相等，一致
		if left != nil && right != nil && left.Val == right.Val {
			return compare(left.Left, right.Right) && compare(left.Right, right.Left)
		}
		// 左右节点可能一边非空，也可能都非空但值不等
		return false
	}
	return compare(root.Left, root.Right)
}

// // 迭代的方式---层序遍历
// // 将单层分成两部分，从左右两端开始向中间步进并比较
// // 并且注意，null 值我们也得放进队列，也得比较，否则 [1,2,2,null,3,null,3] 这种情况显然会错
// func isSymmetric(root *TreeNode) bool {
// 	if root == nil || (root.Left == nil && root.Right == nil) {
// 		return true
// 	}
// 	queue := make([]*TreeNode, 0)
// 	queue = append(queue, root.Left)
// 	queue = append(queue, root.Right)
// 	for len(queue) != 0 {
// 		size := len(queue)
// 		for i := 0; i < size; i++ {
// 			if i < size/2 {
// 				if queue[i] == nil && queue[size-1-i] == nil {
// 					continue
// 				}
// 				if queue[i] == nil || queue[size-1-i] == nil {
// 					return false
// 				}
// 				if queue[i].Val != queue[size-1-i].Val {
// 					return false
// 				}
// 			}
// 			if queue[i] == nil {
// 				continue
// 			}
// 			queue = append(queue, queue[i].Left)
// 			queue = append(queue, queue[i].Right)
// 		}
// 		queue = queue[size:]
// 	}
// 	return true
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
