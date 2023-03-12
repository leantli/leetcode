package main

import "fmt"

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
// 236. 二叉树的最近公共祖先

// 再考虑有没有更方便的方法？做一个后序遍历，去寻找两个节点
// 当找到其中一个就返回对应的节点？挺巧妙的实现
// 当两边都找到了对应的节点时，就返回当前节点，否则就返回找到了的节点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	l, r := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	return r
}

// // 这道题和 235 很类似，但 235 我们能够借助二叉搜索树的性质，方便地去找到两个节点
// // 这里显然不行，我们得确确实实地找到两个节点
// // 这里我们先像 235 的基础解法做一下
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	pPath := getPathOfNode(root, p)
// 	qPath := getPathOfNode(root, q)
// 	res := &TreeNode{Val: -1}
// 	for i := 0; i < len(pPath) && i < len(qPath) && pPath[i] == qPath[i]; i++ {
// 		res = pPath[i]
// 	}
// 	return res
// }

// // 获取达到 target 节点的路径
// func getPathOfNode(root, target *TreeNode) []*TreeNode {
// 	res := make([]*TreeNode, 0)
// 	path := make([]*TreeNode, 0)
// 	var dfs func(root *TreeNode)
// 	dfs = func(root *TreeNode) {
// 		if root == nil || len(res) > 0 {
// 			return
// 		}
// 		path = append(path, root)
// 		if root.Val == target.Val {
// 			res = append(res, append([]*TreeNode{}, path...)...)
// 			return
// 		}
// 		dfs(root.Left)
// 		dfs(root.Right)
// 		path = path[:len(path)-1]
// 	}
// 	dfs(root)
// 	return res
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	res := lowestCommonAncestor(root, root.Left, root.Right)
	fmt.Println(res)
}
