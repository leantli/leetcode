package main

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/
// 235. 二叉搜索树的最近公共祖先

// 正常情况下，我们需要找到 p 和 q 两个节点分别在哪，再停止向下递归，回溯判断
// 某一节点一左一右都为 true，即表示在左右两边各找到了 p 和 q，就是二者的最近公共祖先
// 但是这里我们得注意一下，这是个二叉搜索树，它的值的遍布是有规律的，我们显然能够利用起来
// 利用二叉搜索树的性质判断这两个节点，是否在当前节点的左子树或右子树中
// 如果二者在同一个子树中，说明当前节点必定不是两个节点的最近公共祖先
// 当二者一个在当前节点的左子树中，一个在当前节点的右子树中
// 则说明当前节点就是二者的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}

// // 看了一下官解的各找一边再找最近公共祖先的操作，感觉这个更麻烦更难....
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	pathP := getPath(root, p)
// 	pathQ := getPath(root, q)
// 	// 得到 找两个节点 各自的路径后
// 	// 因为要找最近公共节点，因此找到最后一个他们路径上相等的节点，其为二者的最近公共祖先
// 	var ancestor *TreeNode
// 	for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
// 		ancestor = pathP[i]
// 	}
// 	return ancestor
// }

// // 根据 root 和 二叉搜索树的性质，去找到 target 节点，并记录其路径
// func getPath(root, target *TreeNode) []*TreeNode {
// 	node := root
// 	path := make([]*TreeNode, 0)
// 	for node != target {
// 		path = append(path, node)
// 		if target.Val < node.Val {
// 			node = node.Left
// 		} else {
// 			node = node.Right
// 		}
// 	}
// 	path = append(path, node)
// 	return path
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
