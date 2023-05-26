package main

import "fmt"

// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 54. 二叉搜索树的第k大节点

// 二叉搜索树，性质右大左小
// 采用中序遍历，递归到的第 k 个结点就是第 k 大

func kthLargest(root *TreeNode, k int) int {
	var res int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil || k < 0 {
			return
		}
		dfs(root.Right)
		fmt.Printf("%d, k=%d", root.Val, k)
		k--
		if k == 0 {
			res = root.Val
			return
		}
		dfs(root.Left)
	}
	dfs(root)
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
