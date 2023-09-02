package main

// https://leetcode.cn/problems/check-completeness-of-a-binary-tree/description/
// 958. 二叉树的完全性检验

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 二叉树的完全性检验，显然每一层的节点序号都是从小到大升序
// 我们可以采用层序遍历，只要某一层的从左向右遍历时，节点序号的增加不是+1
// 显然就出现了 null 节点，就不是完全二叉树了，此时直接返回 false
// 全部遍历完，每个相邻节点都是相差1，那么没有问题
// 还有一个点就是，每层刚开始的节点序号，需要是上一层最后一个节点+1
// 这里我们可以使用一个 pre 记录遍历到的上一个节点的序列号，方便比较
func isCompleteTree(root *TreeNode) bool {
	m := make(map[*TreeNode]int) // 记录每个节点对应的序列号
	queue := []*TreeNode{root}
	m[root] = 1
	var pre int // 上一个遍历到的节点的序列号
	for len(queue) > 0 {
		n := len(queue) // 当前层的节点数量
		for i := 0; i < n; i++ {
			// 当前节点的序列号，不等于上一个遍历到的节点的序列号+1
			if m[queue[i]] != pre+1 {
				return false
			}
			pre++
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
				m[queue[i].Left] = m[queue[i]] * 2
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
				m[queue[i].Right] = m[queue[i]]*2 + 1
			}
		}
		queue = queue[n:]
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
