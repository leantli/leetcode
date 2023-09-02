package main

// https://leetcode.cn/problems/maximum-width-of-binary-tree/description/
// 662. 二叉树最大宽度

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 统计所有层的最大宽度，同时，null 节点也计入长度
// 使用 map 记录每个节点的 idx 值，接着正常层序遍历
// 遍历每一层时，对层的首节点和尾节点的 idx 进行计算，获取这一层的宽度
func widthOfBinaryTree(root *TreeNode) int {
	m := make(map[*TreeNode]int)
	queue := []*TreeNode{root}
	m[root] = 0
	var res int
	for len(queue) > 0 {
		n := len(queue)
		res = max(res, m[queue[n-1]]-m[queue[0]]+1)
		for i := 0; i < n; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
				m[queue[i].Left] = m[queue[i]]*2 + 1
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
				m[queue[i].Right] = m[queue[i]]*2 + 2
			}
		}
		queue = queue[n:]
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
